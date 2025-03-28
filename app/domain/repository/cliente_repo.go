package repository

import "context"
import "lanchonete/domain/entities"

type ClienteRepository interface {
	CriarCliente(c context.Context, cliente *entities.Cliente) error
	BuscarCliente(c context.Context, CPF string) (entities.Cliente, error)
}
