package handler

import (
	"encoding/json"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	response "lanchonete/domain/responses"
	"lanchonete/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	PedidoIncluirUseCase         usecases.PedidoIncluirUseCase
	PedidoBuscarPorIdUseCase     usecases.PedidoBuscarPorIdUseCase
	PedidoAtualizarStatusUseCase usecases.PedidoAtualizarStatusUseCase
	ProdutoBuscarPorIdUseCase    usecases.ProdutoBuscaPorIdUseCase
	PedidoListarTodosUseCase     usecases.PedidoListarTodosUseCase
}

func NewPedidoHandler(pedidoIncluirUseCase usecases.PedidoIncluirUseCase,
	pedidoBuscarPorIdUseCase usecases.PedidoBuscarPorIdUseCase,
	pedidoAtualizarStatusUsecase usecases.PedidoAtualizarStatusUseCase,
	produtoBuscarPorIdUseCase usecases.ProdutoBuscaPorIdUseCase,
	pedidoListarTodosUseCase usecases.PedidoListarTodosUseCase) *PedidoHandler {
	return &PedidoHandler{
		PedidoIncluirUseCase:         pedidoIncluirUseCase,
		PedidoBuscarPorIdUseCase:     pedidoBuscarPorIdUseCase,
		PedidoAtualizarStatusUseCase: pedidoAtualizarStatusUsecase,
		ProdutoBuscarPorIdUseCase:    produtoBuscarPorIdUseCase,
		PedidoListarTodosUseCase:     pedidoListarTodosUseCase,
	}
}

// CriarPedido godoc
// @Summary Cria um pedido
// @Description Cria um pedido
// @Tags pedido
// @Router /pedidos [post]
// @Accept  json
// @Produce  json
// @Param pedido body entities.Pedido true "Pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) CriarPedido(r *gin.Context) {
	var pedido entities.Pedido
	err := json.NewDecoder(r.Request.Body).Decode(&pedido)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	for _, produto := range pedido.Produtos {
		_, err = h.ProdutoBuscarPorIdUseCase.Run(r, produto.Identificacao)
		if err != nil {
			r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Produto não Cadastrado!"})
			return
		}
	}
	ped_ret, err := h.PedidoIncluirUseCase.Run(r, pedido.Cliente, pedido.Produtos, pedido.Personalizacao)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido criado com sucesso" + ped_ret.Identificacao,
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags pedido
// @Router /pedidos/{ID} [get]
// @Accept  json
// @Produce  json
// @Param ID path string true "Número do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) BuscarPedido(r *gin.Context) {
	nroPedido := r.Param("nroPedido")
	pedido, err := h.PedidoBuscarPorIdUseCase.Run(r, nroPedido)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, pedido)

}

// BuscarPedido godoc
// @Summary Atualiza um pedido a partir de sua Identificação
// @Description Atualizar um pedido
// @Tags pedido
// @Router /pedidos/{nroPedido}/status/{status} [put]
// @Accept  json
// @Produce  json
// @Param nroPedido path string true "Número do pedido"
// @Param status path string true "Novo Status do pedido"
// @Success 200 {object} entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) AtualizarStatusPedido(r *gin.Context) {
	nroPedido := r.Param("nroPedido")
	status := r.Param("status")
	err := h.PedidoAtualizarStatusUseCase.Run(r, nroPedido, status)
	if err != nil {
		r.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Status do pedido atualizado com sucesso",
	})
}

// ProdutoListarTodos godoc
// @Summary Lista todos os pedidos no banco
// @Description Lista todos os pedidos presentes no banco
// @Tags pedido
// @Router /pedidos/listartodos [POST]
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.Pedido
// @Failure 400 {object} response.ErrorResponse
func (h *PedidoHandler) ListarTodosOsPedidos(r *gin.Context) {
	pedidos, err := h.PedidoListarTodosUseCase.Run(r)
	if err != nil {
		r.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	r.JSON(http.StatusOK, pedidos)
}
