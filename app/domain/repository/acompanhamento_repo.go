package repository

import (
	"context"
	"lanchonete/domain/entities"
)

type AcompanhamentoRepository interface {
	CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedidos(c context.Context, ID string) (entities.Pedido, error)
	AdicionarPedido(c context.Context, acompanhamento *entities.AcompanhamentoPedido, pedido *entities.Pedido) error
	BuscarAcompanhamento(c context.Context, ID string) (*entities.AcompanhamentoPedido, error)
	AtualizarStatusPedido(c context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error
}
