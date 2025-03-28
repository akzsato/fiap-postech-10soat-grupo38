package handler

import (
	"encoding/json"
	"fmt"
	"io"
	_ "lanchonete/docs"
	"lanchonete/domain/entities"
	"lanchonete/domain/responses"
	"lanchonete/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PagamentoHandler struct {
	EnviarPagamentoUseCase usecases.EnviarPagamentoUseCase
	ConfirmarPagamentoUseCase usecases.ConfirmarPagamentoUseCase
}

// EnviarPagamento godoc
// @Summary Envia o pagamento para o webhook	
// @Description Envia o pagamento para o webhook
// @Tags pagamento
// @Router /pagamento [post]
// @Accept  json
// @Produce  json
// @Param pagamento body entities.Pagamento true "Pagamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ph *PagamentoHandler) EnviarPagamento(c *gin.Context) {
	var pagamento entities.Pagamento

	err := c.ShouldBind(&pagamento)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = ph.EnviarPagamentoUseCase.EnviarPagamento(c, &pagamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagamento enviado com sucesso",
	})
}

// ConfirmarPagamento godoc
// @Summary Confirma o pagamento	
// @Description Confirma o pagamento
// @Tags pagamento
// @Router /pagamento/confirmar [post]
// @Accept  json
// @Produce  json
// @Param pagamento body entities.Pagamento true "Pagamento"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
func (ph *PagamentoHandler) ConfirmarPagamento(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Failed to read request body: " + err.Error()})
		return
	}
	defer c.Request.Body.Close()

	fmt.Println("Received webhook data:", string(body))

	var pagamento entities.Pagamento
	if err := json.Unmarshal(body, &pagamento); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Failed to unmarshal request body: " + err.Error()})
		return
	}

	fmt.Printf("Parsed pagamento: %+v\n", pagamento)

	if err := ph.ConfirmarPagamentoUseCase.ConfirmarPagamento(c, &pagamento); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Pagamento confirmado com sucesso",
	})
}
