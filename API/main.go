package main

import (
	"API/DB"
	"API/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//starting gin router
	r := gin.Default()
	//Connection database //Conexão banco de dados
	db := DB.SetupModels()
	//Set on gin context the variable db
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	//Cadastrar um Super/Vilão
	r.POST( "/create/user",controllers.CreateSuperOrVilan)
	//Listar todos os Super's cadastrados
	r.GET(  "/list/all"   ,controllers.ListAllSuperOrVilan)
	//Listar apenas os Super Heróis
	r.GET("/list/super",controllers.ListAllSupers)
	//Listar apenas os Super Vilões
	r.GET("/list/vilan",controllers.ListAllVilans)
	//Buscar por nome
	r.GET(  "/search/name",controllers.SearchByName)
	//Buscar por 'uuid'
	r.GET(  "/search/uuid",controllers.SearchByUuid)
	//Remover o Super
	r.PATCH("/delete/user",controllers.RemoveSuper)


	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}
