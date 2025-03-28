package handler

import (
	"fmt"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	response "lanchonete/domain/responses"
	service "lanchonete/domain/services"
	"lanchonete/domain/usecases"
	"lanchonete/presenters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AcompanhamentoHandler struct {
	AcompanhamentoUseCase        service.AcompanhamentoUseCase
	PedidoAtualizarStatusUseCase usecases.PedidoAtualizarStatusUseCase
}

func NewAcompanhamentoHandler(auc service.AcompanhamentoUseCase, p usecases.PedidoAtualizarStatusUseCase) *AcompanhamentoHandler {
	return &AcompanhamentoHandler{
		AcompanhamentoUseCase:        auc,
		PedidoAtualizarStatusUseCase: p,
	}
}

// CriarAcompanhamento godoc
// @Summary Cria um acompanhamento
// @Description Cria um acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento [post]
// @Accept  json
// @Produce  json
// @Param acompanhamento body entities.AcompanhamentoPedido true "Acompanhamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) CriarAcompanhamento(a *gin.Context) {
	fmt.Print(a)
	var acompanhamento entities.AcompanhamentoPedido

	err := a.ShouldBind(&acompanhamento)
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ah.AcompanhamentoUseCase.CriarAcompanhamento(a, &acompanhamento)
	if err != nil {
		a.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Acompanhamento criado com sucesso",
	})
}

// BuscarPedido godoc
// @Summary Busca um pedido
// @Description Busca um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{ID} [get]
// @Accept  json
// @Produce  json
// @Param ID path string true "ID do pedido"
// @Success 200 {object} entities.AcompanhamentoPedido
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) BuscarPedido(p *gin.Context) {
	ID := p.Param("ID")
	fmt.Print(p.Params)
	pedido, err := ah.AcompanhamentoUseCase.BuscarPedido(p, ID)
	if err != nil {
		p.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	p.JSON(http.StatusOK, pedido)
}

// AdicionarPedido godoc
// @Summary Adiciona um pedido ao acompanhamento
// @Description Adiciona um pedido existente ao acompanhamento de pedidos
// @Tags acompanhamento
// @Router /acompanhamento/{IDAcompanhamento}/{IDPedido} [post]
// @Accept json
// @Produce json
// @Param IDAcompanhamento path string true "ID do acompanhamento"
// @Param IDPedido path string true "ID do pedido"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse "Pedido ou acompanhamento não encontrado"
// @Failure 500 {object} response.ErrorResponse "Erro interno"
func (ah *AcompanhamentoHandler) AdicionarPedido(a *gin.Context) {
	idAcompanhamento := a.Param("IDAcompanhamento")
	idPedido := a.Param("IDPedido")

	pedido := &entities.Pedido{
		Identificacao: idPedido,
	}
	err := ah.AcompanhamentoUseCase.AdicionarPedido(a, idAcompanhamento, pedido)
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pedido adicionado com sucesso",
	})
}

// BuscarAcompanhamento godoc
// @Summary Busca um acompanhamento
// @Description Busca um acompanhamento
// @Tags acompanhamento
// @Router /acompanhamento/show [get]
// @Accept  json
// @Produce  json
// @Param ID path string true "ID do acompanhamento"
// @Success 200 {object} presenters.AcompanhamentoDTO
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) BuscarAcompanhamento(c *gin.Context) {
	acompanhamento, err := ah.AcompanhamentoUseCase.BuscarAcompanhamento(c, "")
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			newAcompanhamento := entities.AcompanhamentoPedido{}
			err = ah.AcompanhamentoUseCase.CriarAcompanhamento(c, &newAcompanhamento)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
				return
			}
			acompanhamento = &newAcompanhamento
		} else {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
			return
		}
	}

	acompanhamentoDto := presenters.NewAcompanhamentoDTO(acompanhamento)
	c.JSON(http.StatusOK, acompanhamentoDto)
}

// AtualizarStatusPedido godoc
// @Summary Atualiza o status de um pedido
// @Description Atualiza o status de um pedido
// @Tags acompanhamento
// @Router /acompanhamento/{IDAcompanhamento}/{IDPedido}/{status} [put]
// @Accept  json
// @Produce  json
// @Param IDAcompanhamento path string true "ID do acompanhamento"
// @Param IDPedido path string true "ID do pedido"
// @Param status path string true "Novo status"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ah *AcompanhamentoHandler) AtualizarStatusPedido(a *gin.Context) {
	idAcompanhamento := a.Param("IDAcompanhamento")
	idPedido := a.Param("IDPedido")
	status := a.Param("status")

	fmt.Printf("Parâmetros recebidos: IDAcompanhamento=%s, IDPedido=%s, status=%s\n",
		idAcompanhamento, idPedido, status)

	err := ah.PedidoAtualizarStatusUseCase.Run(a, idPedido, status)
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ah.AcompanhamentoUseCase.AtualizarStatusPedido(a, idAcompanhamento, idPedido, entities.StatusPedido(status))
	if err != nil {
		a.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	a.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Status atualizado com sucesso",
	})
}

type StatusUpdateRequest struct {
	Status string `json:"status" example:"Em preparação"`
}
