package route

import (
	"lanchonete/bootstrap"
	"lanchonete/infra/database/mongo"

	"github.com/gin-gonic/gin"
)

const collectionCliente = "cliente"
const collectionPagamento = "pagamento"
const CollectionPedido = "pedido"
const CollectionProduto = "produto"
const CollectionAcompanhamento = "acompanhamento"

func Setup(env *bootstrap.Env, db mongo.Database, gin gin.IRouter) {
	router := gin.Group("")
		
	NewDocRouter(router)
	NewClienteRouter(env, db, router)
	NewPagamentoRouter(env, db, router)
	NewPedidoRouter(env, db, router)
	NewAcompanhamentoRouter(env, db, router)
	NewProdutoRouter(env, db, router)
}