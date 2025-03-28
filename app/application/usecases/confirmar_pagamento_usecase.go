package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
)

type confirmarPagamentoUseCase struct {
	pagamentoRepository repository.PagamentoRepository
}

func NewConfirmarPagamentoUseCase(pagamentoRepository repository.PagamentoRepository) usecases.ConfirmarPagamentoUseCase {
	return &confirmarPagamentoUseCase{
		pagamentoRepository: pagamentoRepository,
	}
}

func (uc *confirmarPagamentoUseCase) ConfirmarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	return uc.pagamentoRepository.ConfirmarPagamento(c, pagamento)
}

