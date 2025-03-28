package gateways

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type ProdutoGateway struct {
	database   mongo.Database
	collection string
}

func NewProdutoGateway(db mongo.Database, collection string) repository.ProdutoRepository {
	return &ProdutoGateway{
		database:   db,
		collection: collection,
	}
}

func (pd *ProdutoGateway) AdicionarProduto(c context.Context, produto *entities.Produto) error {
	collection := pd.database.Collection(pd.collection)
	_, err := collection.InsertOne(c, produto)

	return err
}

func (pd *ProdutoGateway) BuscarProdutoPorId(c context.Context, id string) (*entities.Produto, error) {
	collection := pd.database.Collection(pd.collection)

	var produto *entities.Produto

	err := collection.FindOne(c, bson.M{"identificacao": id}).Decode(&produto)
	return produto, err
}

func (pd *ProdutoGateway) ListarTodosOsProdutos(c context.Context) ([]*entities.Produto, error) {
	collection := pd.database.Collection(pd.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

	var produtos []*entities.Produto
	if err = cursor.All(c, &produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}

func (pd *ProdutoGateway) EditarProduto(c context.Context, produto *entities.Produto) error {
	collection := pd.database.Collection(pd.collection)
	filter := bson.M{"identificacao": produto.Identificacao}
	update := bson.M{"$set": bson.M{
		"nome":      produto.Nome,
		"categoria": produto.Categoria,
		"descricao": produto.Descricao,
		"preco":     produto.Preco,
	},
	}
	_, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (pd *ProdutoGateway) RemoverProduto(c context.Context, identificacao string) error {
	collection := pd.database.Collection(pd.collection)
	_, err := collection.DeleteOne(c, bson.M{"identificacao": identificacao})
	return err
}

func (pd *ProdutoGateway) ListarPorCategoria(c context.Context, categoria string) ([]*entities.Produto, error) {
	collection := pd.database.Collection(pd.collection)

	cursor, err := collection.Find(c, bson.M{"categoria": categoria})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

	var produtos []*entities.Produto
	if err = cursor.All(c, &produtos); err != nil {
		return nil, err
	}

	return produtos, nil
}
