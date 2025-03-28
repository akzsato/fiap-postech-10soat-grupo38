package usecases

import (
	"context"
	"fmt"
	"lanchonete/domain/entities"
	"lanchonete/domain/repository"
)

type ProdutoIncluirUseCase struct {
	produtoRepository repository.ProdutoRepository
}

func NewProdutoIncluirUseCase(produtoRepository repository.ProdutoRepository) *ProdutoIncluirUseCase {
	return &ProdutoIncluirUseCase{
		produtoRepository: produtoRepository,
	}
}

func (pd *ProdutoIncluirUseCase) Run(c context.Context, identificacao string, nome string, categoria string, descricao string, preco float32) (*entities.Produto, error) {

	produto, err := entities.ProdutoNew(identificacao, nome, categoria, descricao, preco)

	if err != nil {
		return nil, fmt.Errorf("criação de produto inválida: %w", err)
	}

	err = pd.produtoRepository.AdicionarProduto(c, produto)
	if err != nil {
		return nil, fmt.Errorf("não foi possível criar produto: %w", err)

	}

	return produto, nil
}

// 	"lanchonete/domain/entities"
// )

// type ProdutoIncluirUseCase interface {
// 	Run(c context.Context, identificacao string, nome string, categoria string, descricao string, preco float32) (*entities.Produto, error)
// }
