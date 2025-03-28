package entities

import (
	"errors"
	"strings"
)

type Pagamento struct {
	IdPagamento  string
	Valor string
	Status   string
	DataCriacao string
}

func PagamentoNew(idPagamento string, valor string, status string, dataCriacao string) (*Pagamento, error) {

	if  strings.TrimSpace(idPagamento) == "" || strings.TrimSpace(status) == "" || strings.TrimSpace(dataCriacao) == "" {
		return nil, errors.New("nenhum dos campos podem estar em branco")
	}	

	return &Pagamento{
		IdPagamento:  idPagamento,
		Valor: valor,
		Status: status,	
		DataCriacao: dataCriacao,
	}, nil
}