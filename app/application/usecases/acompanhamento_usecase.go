package usecase

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
	service "lanchonete/domain/services"
)

type acompanhamentoUseCase struct {
	acompanhamentoRepo repository.AcompanhamentoRepository
}

func NewAcompanhamentoUseCase(acompanhamentoRepo repository.AcompanhamentoRepository) service.AcompanhamentoUseCase {
	return &acompanhamentoUseCase{
		acompanhamentoRepo: acompanhamentoRepo,
	}
}

func (uc *acompanhamentoUseCase) CriarAcompanhamento(c context.Context, acompanhamento *entities.AcompanhamentoPedido) error {
	return uc.acompanhamentoRepo.CriarAcompanhamento(c, acompanhamento)
}

func (uc *acompanhamentoUseCase) BuscarPedido(c context.Context, ID string) (entities.Pedido, error) {
	return uc.acompanhamentoRepo.BuscarPedidos(c, ID)
}
func (uc *acompanhamentoUseCase) AdicionarPedido(c context.Context, acompanhamentoID string, pedido *entities.Pedido) error {
	acompanhamento := &entities.AcompanhamentoPedido{ID: acompanhamentoID}
	return uc.acompanhamentoRepo.AdicionarPedido(c, acompanhamento, pedido)
}

func (uc *acompanhamentoUseCase) BuscarAcompanhamento(c context.Context, ID string) (*entities.AcompanhamentoPedido, error) {
	return uc.acompanhamentoRepo.BuscarAcompanhamento(c, ID)
}

func (uc *acompanhamentoUseCase) AtualizarStatusPedido(c context.Context, acompanhamentoID string, identificacao string, novoStatus entities.StatusPedido) error {
	fmt.Printf("UseCase: Atualizando pedido %s no acompanhamento %s para status %s\n",
		identificacao, acompanhamentoID, novoStatus)

	return uc.acompanhamentoRepo.AtualizarStatusPedido(c, acompanhamentoID, identificacao, novoStatus)
}
