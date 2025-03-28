package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type ClienteUseCase interface {
	CriarCliente(c context.Context, cliente *entities.Cliente) error
	BuscarCliente(c context.Context, cpf string) (entities.Cliente, error)
}