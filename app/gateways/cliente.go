package gateways

import (
	"context"
	"fmt"

	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type clienteRepository struct {
	database mongo.Database
	collection string
}

func NewClienteRepository(db mongo.Database, collection string) repository.ClienteRepository {
	return &clienteRepository{
		database: db,
		collection: collection,
	}
}

func (cr *clienteRepository) CriarCliente(c context.Context, cliente *entities.Cliente) error {
	collection := cr.database.Collection(cr.collection)
	fmt.Println("COLLECTION", collection)
	_, err := collection.InsertOne(c, cliente)

	return err
}

func (cr *clienteRepository) BuscarCliente(c context.Context, CPF string) (entities.Cliente, error) {
	collection := cr.database.Collection(cr.collection)

	var cliente entities.Cliente

	err := collection.FindOne(c, bson.M{"cpf": CPF}).Decode(&cliente)
	return cliente, err
}