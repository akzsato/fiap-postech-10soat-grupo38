package gateways

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type PedidoGateway struct {
	database   mongo.Database
	collection string
}

func NewPedidoGateway(db mongo.Database, collection string) repository.PedidoRepository {
	return &PedidoGateway{
		database:   db,
		collection: collection,
	}
}

func (pg *PedidoGateway) CriarPedido(c context.Context, pedido *entities.Pedido) error {
	collection := pg.database.Collection(pg.collection)

	_, err := collection.InsertOne(c, pedido)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PedidoGateway) BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error) {
	collection := pg.database.Collection(pg.collection)

	var pedido entities.Pedido

	err := collection.FindOne(c, bson.M{"identificacao": identificacao}).Decode(&pedido)
	return &pedido, err
}

func (pg *PedidoGateway) AtualizarStatusPedido(c context.Context, identificacao string, status string, ultimaAtualizacao string) error {
	collection := pg.database.Collection(pg.collection)

	filter := bson.M{"identificacao": identificacao}
	update := bson.M{"$set": bson.M{"status": status,
		"UltimaAtualizacao": ultimaAtualizacao,
	},
	}

	_, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PedidoGateway) ListarTodosOsPedidos(c context.Context) ([]*entities.Pedido, error) {
	collection := pg.database.Collection(pg.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

	var pedidos []*entities.Pedido
	if err = cursor.All(c, &pedidos); err != nil {
		return nil, err
	}

	return pedidos, nil
}
