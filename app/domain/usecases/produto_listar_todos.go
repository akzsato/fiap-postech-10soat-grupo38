package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type ProdutoListarTodosUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func NewProdutoListarTodosUseCase(produtoRepo repository.ProdutoRepository) *ProdutoListarTodosUseCase {
	return &ProdutoListarTodosUseCase{
		produtoRepo: produtoRepo,
	}
}

func (pd *ProdutoListarTodosUseCase) Run(c context.Context) ([]*entities.Produto, error) {
	produtos, err := pd.produtoRepo.ListarTodosOsProdutos(c)
	if err != nil {
		return nil, fmt.Errorf("não foi possível listar produtos: %w", err)
	}
	return produtos, nil
}
