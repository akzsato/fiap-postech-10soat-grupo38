package entities

import (
	"sort"
	"sync"
	"time"
)

type FilaPedidos struct {
	Pedidos []Pedido    `bson:"pedidos" json:"pedidos"`
	mu      *sync.Mutex `bson:"-" json:"-"`
}

func NewFilaPedidos() FilaPedidos {
	return FilaPedidos{
		Pedidos: make([]Pedido, 0),
		mu:      &sync.Mutex{},
	}
}

func (f *FilaPedidos) EnsureMutex() {
	if f.mu == nil {
		f.mu = &sync.Mutex{}
	}
}

func (f *FilaPedidos) Enfileirar(p Pedido) {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()

	f.Pedidos = append(f.Pedidos, p)

	sort.Slice(f.Pedidos, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, f.Pedidos[i].TimeStamp)
		timeJ, _ := time.Parse(time.RFC3339, f.Pedidos[j].TimeStamp)
		return timeI.Before(timeJ)
	})
}

func (f *FilaPedidos) Desenfileirar() (Pedido, bool) {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.Pedidos) == 0 {
		return Pedido{}, false
	}

	pedido := f.Pedidos[0]
	f.Pedidos = f.Pedidos[1:]
	return pedido, true
}

func (f *FilaPedidos) RemoverPedido(identificacao string) bool {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()

	for i, pedido := range f.Pedidos {
		if pedido.Identificacao == identificacao {
			f.Pedidos = append(f.Pedidos[:i], f.Pedidos[i+1:]...)
			return true
		}
	}
	return false
}

func (f *FilaPedidos) Listar() []Pedido {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()

	result := make([]Pedido, len(f.Pedidos))
	copy(result, f.Pedidos)
	return result
}

func (f *FilaPedidos) Tamanho() int {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()

	return len(f.Pedidos)
}

func (f *FilaPedidos) IsEmpty() bool {
	f.EnsureMutex()
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.Pedidos) == 0
}
