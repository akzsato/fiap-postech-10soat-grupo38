package usecases

import (
	"context"
	"lanchonete/domain/entities"
)

type EnviarPagamentoUseCase interface {
	EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error
}