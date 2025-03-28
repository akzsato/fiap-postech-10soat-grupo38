package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type ProdutoBuscaPorIdUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func NewProdutoBuscaPorIdUseCase(produtoRepo repository.ProdutoRepository) *ProdutoBuscaPorIdUseCase {
	return &ProdutoBuscaPorIdUseCase{
		produtoRepo: produtoRepo,
	}
}

func (pd *ProdutoBuscaPorIdUseCase) Run(c context.Context, id string) (*entities.Produto, error) {
	produto, err := pd.produtoRepo.BuscarProdutoPorId(c, id)

	if err != nil {
		return nil, fmt.Errorf("não foi possível buscar produto: %w", err)
	}
	return produto, nil
}
