package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type ProdutoListarPorCategoriaUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func NewProdutoListarPorCategoriaUseCase(produtoRepo repository.ProdutoRepository) *ProdutoListarPorCategoriaUseCase {
	return &ProdutoListarPorCategoriaUseCase{
		produtoRepo: produtoRepo,
	}
}

func (pd *ProdutoListarPorCategoriaUseCase) Run(c context.Context, categoria string) ([]*entities.Produto, error) {
	produtos, err := pd.produtoRepo.ListarPorCategoria(c, categoria)
	if err != nil {
		return nil, fmt.Errorf("não foi possível listar produtos: %w", err)
	}
	return produtos, nil
}
