package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/repository"
)

type ProdutoRemoverUseCase struct {
	produtoGateway repository.ProdutoRepository
}

func NewProdutoRemoverUseCase(produtoGateway repository.ProdutoRepository) *ProdutoRemoverUseCase {
	return &ProdutoRemoverUseCase{
		produtoGateway: produtoGateway,
	}
}

func (pruc *ProdutoRemoverUseCase) Run(c context.Context, identificacao string) error {
	_, err := pruc.produtoGateway.BuscarProdutoPorId(c, identificacao)
	if err != nil {
		return fmt.Errorf("produto não existe no banco de dados: %w", err)
	}

	err = pruc.produtoGateway.RemoverProduto(c, identificacao)
	if err != nil {
		return fmt.Errorf("não foi possível remover o produto: %w", err)
	}

	return nil
}
