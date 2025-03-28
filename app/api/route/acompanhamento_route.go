package route

import (
	handler "lanchonete/api/handlers"
	usecase "lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/domain/usecases"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"fmt"

	"github.com/gin-gonic/gin"
)

func NewAcompanhamentoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	ar := gateways.NewAcompanhamentoRepository(db, CollectionAcompanhamento)
	pr := gateways.NewPedidoGateway(db, CollectionPedido)
	auc := &handler.AcompanhamentoHandler{
		AcompanhamentoUseCase:        usecase.NewAcompanhamentoUseCase(ar),
		PedidoAtualizarStatusUseCase: *usecases.NewPedidoAtualizarStatusUseCase(pr),
	}

	fmt.Printf("Registrando rotas do acompanhamento\n")

	router.POST("/acompanhamento", auc.CriarAcompanhamento)
	router.GET("/acompanhamento/show", auc.BuscarAcompanhamento)
	router.GET("/acompanhamento/:ID", auc.BuscarPedido)
	router.POST("/acompanhamento/:IDAcompanhamento/:IDPedido", auc.AdicionarPedido)
	router.PUT("acompanhamento/:IDAcompanhamento/:IDPedido/:status", auc.AtualizarStatusPedido)
}
