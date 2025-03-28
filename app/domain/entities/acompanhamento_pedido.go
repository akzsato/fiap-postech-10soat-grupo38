package entities

import (
	"time"
)

const (
	DefaultPreparationTime = 15 * time.Minute
)

type AcompanhamentoPedido struct {
	ID               string
	Pedidos          FilaPedidos
	TempoEstimado    time.Duration
	UltimaAtualizacao string
}

func NewAcompanhamentoPedido() *AcompanhamentoPedido {
	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	return &AcompanhamentoPedido{
		Pedidos:           NewFilaPedidos(),
		TempoEstimado:     DefaultPreparationTime,
		UltimaAtualizacao: agora,
	}
}

func (a *AcompanhamentoPedido) AdicionarPedido(pedido Pedido) {
	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	if pedido.Status == Recebido {
		a.Pedidos.Enfileirar(pedido)
		a.UltimaAtualizacao = agora
	}
}

func (a *AcompanhamentoPedido) AtualizarStatusPedido(identificacao string, novoStatus StatusPedido) bool {
	local, _ := time.LoadLocation("America/Sao_Paulo")
	agora := time.Now().In(local).Format("2006-01-02 15:04:05")
	pedidos := a.Pedidos.Listar()
	for i, pedido := range pedidos {
		if pedido.Identificacao == identificacao {
			pedido.Status = novoStatus
			pedido.UltimaAtualizacao = agora

			if novoStatus == Finalizado {
				a.Pedidos = a.removerPedidoNaPosicao(i)
			} else {
				pedidos[i] = pedido
			}

			a.UltimaAtualizacao = agora
			return true
		}
	}
	return false
}

func (a *AcompanhamentoPedido) removerPedidoNaPosicao(posicao int) FilaPedidos {
	pedidos := a.Pedidos.Listar()
	novosPedidos := append(pedidos[:posicao], pedidos[posicao+1:]...)

	novaFila := NewFilaPedidos()
	for _, p := range novosPedidos {
		novaFila.Enfileirar(p)
	}

	return novaFila
}
