package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type PedidoListarTodosUseCase struct {
	pedidoRepo repository.PedidoRepository
}

func NewPedidoListarTodosUseCase(pedidoRepo repository.PedidoRepository) *PedidoListarTodosUseCase {
	return &PedidoListarTodosUseCase{
		pedidoRepo: pedidoRepo,
	}
}

func (pd *PedidoListarTodosUseCase) Run(c context.Context) ([]*entities.Pedido, error) {
	pedidos, err := pd.pedidoRepo.ListarTodosOsPedidos(c)
	if err != nil {
		return nil, fmt.Errorf("não foi possível listar pedidos: %w", err)
	}
	return pedidos, nil
}
