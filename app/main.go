package main

import (
	route "lanchonete/api/route"
	"lanchonete/bootstrap"
	_ "lanchonete/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Lanchonete API - Tech Challenge 2
// @version 1.0
// @description API para o Tech Challenge 2 da FIAP - SOAT

// @host localhost:8080
// @BasePath /
//
//go:generate go run github.com/swaggo/swag/cmd/swag@latest init
func main() {
	gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.Default()

	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	route.Setup(env, db, ginEngine)

	ginEngine.Run(env.ServerAddress)
}
