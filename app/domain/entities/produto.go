package entities

import (
	"errors"
	"strings"
)

type CatProduto string

const (
	Lanche         CatProduto = "Lanche"
	Acompanhamento CatProduto = "Acompanhamento"
	Bebida         CatProduto = "Bebida"
	Sobremesa      CatProduto = "Sobremesa"
)

type Produto struct {
	Identificacao string
	Nome          string
	Categoria     CatProduto
	Descricao     string
	Preco         float32
}

func ProdutoNew(identificacao string, nome string, categoria string, descricao string, preco float32) (*Produto, error) {
	if strings.TrimSpace(identificacao) == "" || strings.TrimSpace(nome) == "" || preco <= 0 || strings.TrimSpace(categoria) == "" {
		return nil, errors.New("todos os campos são obrigatórios e o preço maior que zero")
	}

	var cat_prod CatProduto

	switch CatProduto(categoria) {
	case Lanche, Acompanhamento, Bebida, Sobremesa:
		cat_prod = CatProduto(categoria)
	default:
		return nil, errors.New("categoria inválida")
	}

	return &Produto{
		Identificacao: identificacao,
		Nome:          nome,
		Categoria:     cat_prod,
		Descricao:     descricao,
		Preco:         preco,
	}, nil
}
