package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionAcompanhamento = "acompanhamento"
)

type AcompanhamentoUseCase interface {
	CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error
	BuscarPedido(c context.Context, nroPedido string) (entities.Pedido, error)
	AdicionarPedido(c context.Context, acompanhamentoID string, pedido *entities.Pedido) error
	BuscarAcompanhamento(c context.Context, nroPedido string) (*entities.AcompanhamentoPedido, error)
	AtualizarStatusPedido(c context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error
}
