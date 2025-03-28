package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type ConfirmarPagamentoUseCase interface {
	ConfirmarPagamento(C context.Context, pagamento *entities.Pagamento) error
}