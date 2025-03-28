package gateways

import (
	"context"
	"fmt"
	"time"

	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type acompanhamentoRepository struct {
	database          mongo.Database
	collection        string
	pedidosCollection string
}

func NewAcompanhamentoRepository(db mongo.Database, collection string) repository.AcompanhamentoRepository {
	return &acompanhamentoRepository{
		database:          db,
		collection:        collection,
		pedidosCollection: "pedido",
	}
}

func (ar *acompanhamentoRepository) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	collection := ar.database.Collection(ar.collection)
	_, err := collection.InsertOne(c, acompanhamento)
	return err
}

func (ar *acompanhamentoRepository) BuscarPedidos(c context.Context, ID string) (entities.Pedido, error) {
	collection := ar.database.Collection(ar.pedidosCollection)
	fmt.Printf("Collection name: %s\n", ar.pedidosCollection)

	fmt.Printf("Buscando pedido com ID: %s\n", ID)

	filter := bson.M{
		"$or": []bson.M{
			{"identificacao": ID},
		},
	}

	var pedido entities.Pedido
	err := collection.FindOne(c, filter).Decode(&pedido)

	if err != nil {
		fmt.Printf("Erro ao buscar pedido: %v\n", err)
		return entities.Pedido{}, fmt.Errorf("pedido não encontrado: %v", err)
	}

	fmt.Printf("Pedido encontrado: %+v\n", pedido)
	return pedido, nil
}

func (ar *acompanhamentoRepository) AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, p *entities.Pedido) error {
	fmt.Printf("Tentando adicionar pedido %s ao acompanhamento %s\n", p.Identificacao, acompanhamento.ID)

	// Buscar o pedido
	pedido, err := ar.BuscarPedidos(c, p.Identificacao)
	if err != nil {
		return fmt.Errorf("erro ao buscar pedido: %v", err)
	}
	fmt.Printf("Pedido encontrado: %+v\n", pedido)

	// Listar todos os documentos na coleção
	cursor, err := ar.database.Collection(ar.collection).Find(c, bson.M{})
	if err != nil {
		return fmt.Errorf("erro ao listar documentos: %v", err)
	}
	defer cursor.Close(c)

	var documentos []bson.M
	if err = cursor.All(c, &documentos); err != nil {
		return fmt.Errorf("erro ao decodificar documentos: %v", err)
	}

	fmt.Printf("Encontrados %d documentos na coleção\n", len(documentos))

	// Procurar o documento com o ID correto
	var documentoAlvo bson.M
	var documentoID interface{}
	var idField string

	for _, doc := range documentos {
		// Verificar todos os campos possíveis de ID
		for _, field := range []string{"id", "ID"} {
			if id, exists := doc[field]; exists {
				idStr := fmt.Sprintf("%v", id) // Converter para string para comparação
				fmt.Printf("Documento com %s=%v\n", field, idStr)

				if idStr == acompanhamento.ID {
					documentoAlvo = doc
					documentoID = id
					idField = field
					fmt.Printf("Documento alvo encontrado com %s=%v\n", field, id)
					break
				}
			}
		}
		if documentoAlvo != nil {
			break
		}
	}

	if documentoAlvo == nil {
		return fmt.Errorf("acompanhamento com ID %s não encontrado em nenhum documento", acompanhamento.ID)
	}

	// Construir o filtro com o campo de ID correto
	filter := bson.M{idField: documentoID}
	fmt.Printf("Usando filtro para atualização: %+v\n", filter)

	// Verificar a estrutura atual do documento
	var acompanhamentoAtual entities.AcompanhamentoPedido
	err = ar.database.Collection(ar.collection).FindOne(c, filter).Decode(&acompanhamentoAtual)
	if err != nil {
		return fmt.Errorf("erro ao decodificar acompanhamento: %v", err)
	}

	// Adicionar o pedido à fila
	acompanhamentoAtual.Pedidos.Enfileirar(pedido)

	// Obter a lista de pedidos atualizada
	pedidosAtualizados := acompanhamentoAtual.Pedidos.Listar()
	fmt.Printf("Pedidos após adicionar: %+v\n", pedidosAtualizados)

	// Atualizar o documento
	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	update := bson.M{
		"$set": bson.M{
			"pedidos":           bson.M{"pedidos": pedidosAtualizados},
			"ultimaAtualizacao": agora,
		},
	}

	result, err := ar.database.Collection(ar.collection).UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("erro ao atualizar acompanhamento: %v", err)
	}

	fmt.Printf("Resultado da atualização: ModifiedCount=%d, MatchedCount=%d\n",
		result.ModifiedCount, result.MatchedCount)

	if result.ModifiedCount == 0 {
		return fmt.Errorf("nenhum documento foi atualizado")
	}

	return nil
}

func (ar *acompanhamentoRepository) BuscarAcompanhamento(ctx context.Context, ID string) (*entities.AcompanhamentoPedido, error) {
	var acompanhamento *entities.AcompanhamentoPedido

	fmt.Printf("Buscando acompanhamento na collection: %s\n", ar.collection)

	filter := bson.M{}
	if ID != "" {
		filter["id"] = ID
	}

	fmt.Printf("Usando filter: %+v\n", filter)

	err := ar.database.Collection(ar.collection).FindOne(ctx, filter).Decode(&acompanhamento)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar acompanhamento: %v", err)
	}

	fmt.Printf("Acompanhamento encontrado: %+v\n", acompanhamento)
	fmt.Printf("Pedidos na fila: %+v\n", acompanhamento.Pedidos.Listar())
	fmt.Printf("Número de pedidos na fila: %d\n", acompanhamento.Pedidos.Tamanho())

	acompanhamento.Pedidos.EnsureMutex()
	return acompanhamento, nil
}

func (ar *acompanhamentoRepository) AtualizarStatusPedido(ctx context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error {
	fmt.Printf("Repository: Atualizando pedido %s no acompanhamento %s para status %s\n",
		identificacao, acompanhamentoID, novoStatus)

	cursor, err := ar.database.Collection(ar.collection).Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("erro ao listar documentos: %v", err)
	}
	defer cursor.Close(ctx)

	var documentos []bson.M
	if err = cursor.All(ctx, &documentos); err != nil {
		return fmt.Errorf("erro ao decodificar documentos: %v", err)
	}

	fmt.Printf("Encontrados %d documentos na coleção\n", len(documentos))

	var documentoAlvo bson.M
	var documentoID interface{}
	var idField string

	for _, doc := range documentos {
		for _, field := range []string{"id", "_id", "ID"} {
			if id, exists := doc[field]; exists {
				idStr := fmt.Sprintf("%v", id)
				fmt.Printf("Documento com %s=%v\n", field, idStr)

				if idStr == acompanhamentoID {
					documentoAlvo = doc
					documentoID = id
					idField = field
					fmt.Printf("Documento alvo encontrado com %s=%v\n", field, id)
					break
				}
			}
		}
		if documentoAlvo != nil {
			break
		}
	}

	if documentoAlvo == nil {
		return fmt.Errorf("acompanhamento com ID %s não encontrado em nenhum documento", acompanhamentoID)
	}

	filter := bson.M{idField: documentoID}
	fmt.Printf("Usando filtro para atualização: %+v\n", filter)

	if pedidos, ok := documentoAlvo["pedidos"].(bson.M); ok {
		if pedidosArray, ok := pedidos["pedidos"].(bson.A); ok {
			fmt.Printf("Estrutura de pedidos encontrada com %d pedidos\n", len(pedidosArray))

			for i, p := range pedidosArray {
				if pedido, ok := p.(bson.M); ok {
					fmt.Printf("Pedido %d: ID=%v, Status=%v\n", i, pedido["identificacao"], pedido["status"])
				}
			}
		}
	}

	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	update := bson.M{
		"$set": bson.M{
			"pedidos.pedidos.$[elem].status":            string(novoStatus),
			"pedidos.pedidos.$[elem].ultimaatualizacao": agora,
			"ultimaAtualizacao":                         agora,
		},
	}

	// Remover se finalizado
	if novoStatus == entities.Finalizado {
		removeUpdate := bson.M{
			"$pull": bson.M{
				"pedidos.pedidos": bson.M{"identificacao": identificacao},
			},
		}
		_, err = ar.database.Collection(ar.collection).UpdateOne(ctx, filter, removeUpdate)
		if err != nil {
			return fmt.Errorf("erro ao remover pedido finalizado: %v", err)
		}
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.identificacao": identificacao},
		},
	})

	result, err := ar.database.Collection(ar.collection).UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return fmt.Errorf("erro ao atualizar pedido: %v", err)
	}

	fmt.Printf("Resultado da atualização: ModifiedCount=%d, MatchedCount=%d\n",
		result.ModifiedCount, result.MatchedCount)

	if result.ModifiedCount == 0 {
		return fmt.Errorf("pedido %s não encontrado no acompanhamento %s", identificacao, acompanhamentoID)
	}

	return nil
}
