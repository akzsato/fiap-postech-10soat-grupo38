package repository

import (
	"context"
	"lanchonete/domain/entities"
)

type PedidoRepository interface {
	CriarPedido(c context.Context, pedido *entities.Pedido) error
	BuscarPedido(c context.Context, identificacao string) (*entities.Pedido, error)
	AtualizarStatusPedido(c context.Context, Identificacao string, status string, UltimaAtualizacao string) error
	ListarTodosOsPedidos(c context.Context) ([]*entities.Pedido, error)
}
