package usecases

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoAtualizarStatusUseCase struct {
	pedidoGateway repository.PedidoRepository
}

func NewPedidoAtualizarStatusUseCase(pedidoGateway repository.PedidoRepository) *PedidoAtualizarStatusUseCase {
	return &PedidoAtualizarStatusUseCase{
		pedidoGateway: pedidoGateway,
	}
}

func (pduc *PedidoAtualizarStatusUseCase) Run(c context.Context, identificacao string, status string) error {

	pedido, err := pduc.pedidoGateway.BuscarPedido(c, identificacao)
	if err != nil {
		return err
	}

	timeStamp, err := pedido.UpdateStatus(entities.StatusPedido(status))
	if err != nil {
		return err
	}
	err = pduc.pedidoGateway.AtualizarStatusPedido(c, identificacao, status, timeStamp)
	if err != nil {
		return err
	}
	return nil
}
