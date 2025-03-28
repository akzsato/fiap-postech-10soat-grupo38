package service

import (
	"context"
	"lanchonete/domain/entities"
)

const (
	CollectionProduto = "produto"
)

type ProdutoUseCase interface {
	AdicionarProduto(c context.Context, produto *entities.Produto) error
	BuscarProdutoPorId(c context.Context, id string) (entities.Produto, error)
}
