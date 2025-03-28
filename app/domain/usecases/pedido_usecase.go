package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type PedidoUseCase interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, IdDoPedido string) (*entities.Pedido, error)
}
