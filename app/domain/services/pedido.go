package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionPedido = "pedido"
)

type PedidoUseCase interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, identificacao string) (entities.Pedido, error)
}
