package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type StatusPedido string

const (
	Pendente     StatusPedido = "Pendente"
	Recebido     StatusPedido = "Recebido"
	EmPreparacao StatusPedido = "Em preparação"
	Pronto       StatusPedido = "Pronto"
	Finalizado   StatusPedido = "Finalizado"
)

type Pedido struct {
	Identificacao     string			`json:"identificacao"`
	Cliente			  Cliente			`json:"cliente"`	
	Produtos          []Produto			`json:"produtos"`
	Personalizacao	  string			`json:"personalizacao"`
	Status            StatusPedido		`json:"status"`
	StatusPagamento   string			`json:"statusPagamento"`
	TimeStamp         string			`json:"timeStamp" format:"date-time"`
	UltimaAtualizacao string			`json:"ultimaAtualizacao" format:"date-time"`
	Total             float32			`json:"total"`
}

func PedidoNew(cliente Cliente, produtos []Produto, personalizacao string) (*Pedido, error) {
	id := uuid.New()
	id_string := id.String()
	if len(produtos) == 0 {
		return nil, errors.New("o pedido precisa ter ao menos um produto")
	}

	temLanche := false
	for _, produto := range produtos {
		if produto.Categoria == Lanche {
			temLanche = true
			break
		}
	}
	if !temLanche {
		return nil, errors.New("o pedido precisa ter ao menos um lanche")
	}

	total := float32(0)
	for _, produto := range produtos {
		total += produto.Preco
	}

	local, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, errors.New("erro ao carregar localização")
	}
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")

	return &Pedido{
		Identificacao:     id_string,
		Cliente:           cliente,
		Produtos:          produtos,
		Personalizacao:    personalizacao,
		Status:            Pendente,
		StatusPagamento:   "Pendente",
		TimeStamp:         agora,
		UltimaAtualizacao: agora,
		Total:             total,
	}, nil
}

func (p *Pedido) UpdateStatus(status StatusPedido) (ultimaAtualizacao string, err error) {
	switch status {
	case Recebido, EmPreparacao, Pronto, Finalizado:
		p.Status = status
	default:
		return "", errors.New("status inválido")
	}
	local, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return "", errors.New("erro ao carregar localização")
	}
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	p.UltimaAtualizacao = agora

	return p.UltimaAtualizacao, nil
}
