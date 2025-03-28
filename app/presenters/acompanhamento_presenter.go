package presenters

import (
	"lanchonete/domain/entities"
	"time"
)

type PedidoDTO struct {
	ID            string                `json:"id"`
	Identificacao string                `json:"identificacao"`
	Status        entities.StatusPedido `json:"status"`
	TempoEstimado time.Duration         `json:"tempoEstimado"`
}

type AcompanhamentoDTO struct {
	ID            string      `json:"id"`
	Pedidos       []PedidoDTO `json:"pedidos"`
	TempoEstimado int         `json:"tempoEstimado"` // in minutes
}

func NewAcompanhamentoDTO(a *entities.AcompanhamentoPedido) *AcompanhamentoDTO {
	pedidos := make([]PedidoDTO, 0)
	for _, p := range a.Pedidos.Listar() {
		pedidos = append(pedidos, PedidoDTO{
			Identificacao: p.Identificacao,
			Status:        p.Status,
			TempoEstimado: time.Duration(900),
		})
	}
	return &AcompanhamentoDTO{
		ID:            a.ID,
		Pedidos:       pedidos,
		TempoEstimado: int(a.TempoEstimado.Minutes()),
	}
}
