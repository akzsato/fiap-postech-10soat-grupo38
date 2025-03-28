package route

import (
	handler "lanchonete/api/handlers"
	"lanchonete/domain/usecases"

	//usecase "lanchonete/application/usecases"
	"lanchonete/bootstrap"
	"lanchonete/gateways"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

func NewProdutoRouter(env *bootstrap.Env, db mongo.Database, router *gin.RouterGroup) {

	pr := gateways.NewProdutoGateway(db, CollectionProduto)
	pc := &handler.ProdutoHandler{
		ProdutoIncluirUseCase:            *usecases.NewProdutoIncluirUseCase(pr),
		ProdutoBuscarPorIdUseCase:        *usecases.NewProdutoBuscaPorIdUseCase(pr),
		ProdutoListarTodosUseCase:        *usecases.NewProdutoListarTodosUseCase(pr),
		ProdutoEditarUseCase:             *usecases.NewProdutoEditarUseCase(pr),
		ProdutoRemoverUseCase:            *usecases.NewProdutoRemoverUseCase(pr),
		ProdutoListarPorCategoriaUseCase: *usecases.NewProdutoListarPorCategoriaUseCase(pr),
	}

	router.POST("/produto", pc.ProdutoIncluir)
	router.GET("/produto/:id", pc.ProdutoBuscarPorId)
	router.GET("/produtos", pc.ProdutoListarTodos)
	router.GET("/produtos/:categoria", pc.ProdutoListarPorCategoria)
	router.POST("/produto/editar", pc.ProdutoEditar)
	router.DELETE("/produto/:id", pc.ProdutoRemover)
}
