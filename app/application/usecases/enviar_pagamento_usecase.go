package usecase

import (
	"context"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	"lanchonete/domain/usecases"
)

type enviarPagamentoUseCase struct {
	pagamentoRepository repository.PagamentoRepository
}

func NewEnviarPagamentoUseCase(pagamentoRepository repository.PagamentoRepository) usecases.EnviarPagamentoUseCase {
	return &enviarPagamentoUseCase{
		pagamentoRepository: pagamentoRepository,
	}
}

func (uc *enviarPagamentoUseCase) EnviarPagamento(c context.Context, pagamento *entities.Pagamento) error {
	return uc.pagamentoRepository.EnviarPagamento(c, pagamento)
}
