package main

import (
	"API/DB"
	"API/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := DB.SetupModels()
	//Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})
	r.POST("/create", controllers.CreateSuperOrVilan)
	r.GET("/list/all", controllers.ListAllSuperOrVilan)
	r.GET("/search/name",controllers.SearchByName)

	//Cadastrar um Super/Vilão
	//Listar todos os Super's cadastrados
	//Listar apenas os Super Heróis
	//Listar apenas os Super Vilões
	//Buscar por nome
	//Buscar por 'uuid'
	//Remover o Super

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}
