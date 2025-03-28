package handler

import (
	"net/http"
	_"lanchonete/docs"
	"lanchonete/domain/usecases"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	ClienteUseCase usecases.ClienteUseCase
}

// CriarCliente godoc
// @Summary Cria um cliente	
// @Description Cria um cliente
// @Tags cliente
// @Router /cliente [post]
// @Accept  json
// @Produce  json
// @Param cliente body entities.Cliente true "Cliente"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (cc *ClienteHandler) CriarCliente(c *gin.Context) {
	var cliente entities.Cliente

	err := c.ShouldBind(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.ClienteUseCase.CriarCliente(c, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Cliente criado com sucesso",
	})
}

// BuscarCliente godoc
// @Summary Busca um cliente
// @Description Busca um cliente
// @Tags cliente
// @Router /cliente/{CPF} [get]
// @Accept  json
// @Produce  json
// @Param CPF path string true "CPF do cliente"
// @Success 200 {object} entities.Cliente
// @Failure 400 {object} response.ErrorResponse
func (cc *ClienteHandler) BuscarCliente(c *gin.Context) {
	CPF := c.Param("CPF")
	cliente, err := cc.ClienteUseCase.BuscarCliente(c, CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}