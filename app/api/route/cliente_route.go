package route

import (
	"lanchonete/api/handlers"
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"
	"lanchonete/gateways"
	"lanchonete/application/usecases"

	"github.com/gin-gonic/gin"
)


func NewClienteRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {
	cr := gateways.NewClienteRepository(db, collectionCliente)
	cc := &handler.ClienteHandler{
		ClienteUseCase: usecase.NewClienteUseCase(cr),
	}

	router.GET("/cliente/:CPF", cc.BuscarCliente)
	router.POST("/cliente", cc.CriarCliente)
}