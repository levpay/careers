package controllers

import (
	"API/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

//Cadastrar um Super/Vilão
func CreateSuperOrVilan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.SuperOrVilanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Cada super deve ser cadastrado somente a partir do seu name.
	if db.Where("name", input.Name).RecordNotFound() {
		superOrVilan := models.SuperOrVilan{
			Name:         input.Name,
			FullName:     input.FullName,
			Intelligence: input.Intelligence,
			Power:        input.Power,
			Occupation:   input.Occupation,
			Image:        input.Image,
		}
		if err := db.Create(&superOrVilan).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{"success": "Super or Vilan criado !"})
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": "registro existente",
	})

}
//Listar todos os Super's cadastrados
func ListAllSuperOrVilan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var supers []models.SuperOrVilan
	if err := db.Find(&supers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": supers,
	})
	return
}
//Buscar por nome
func SearchByName(c *gin.Context){
	var superDatabase models.SuperOrVilan
	if err:=dataBaseSearch("name",c.Query("name"),c,&superDatabase);err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	var superSearch models.SuperOrVilanSearch
	population(superDatabase,&superSearch)
	c.JSON(http.StatusOK,gin.H{"success":superSearch})
}
//Buscar por 'uuid'
func SearchByUuid(c *gin.Context){
	var superDatabase models.SuperOrVilan
	if err:=dataBaseSearch("name",c.Query("uuid"),c,&superDatabase);err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	var superSearch models.SuperOrVilanSearch
	population(superDatabase,&superSearch)
	c.JSON(http.StatusOK,gin.H{"success":superSearch})
}

//FieldOfSearch
func dataBaseSearch(fieldOfSearch string,Query string,c *gin.Context, superDatabase*models.SuperOrVilan) error{
	db := c.MustGet("db").(*gorm.DB)
	fieldOfSearch=strings.ToLower(fieldOfSearch)
	if fieldOfSearch != "name" || fieldOfSearch !="uuid"{
		return errors.New("parametro invalido")
	}
	//Recuperando os dados do BD
	if err:=db.Where(""+ fieldOfSearch+" = ?",Query).Find(superDatabase).Error;err!=nil{
		return  errors.New(err.Error())
	}
	return nil
}
//esta função popula e devolve para quem chamou
func population(superDatabase models.SuperOrVilan,superSearch *models.SuperOrVilanSearch) {
	superSearch.Uuid=             superDatabase.Uuid             //uuid
	superSearch.Name=             superDatabase.Name             //name
	superSearch.FullName=         superDatabase.FullName         //full name
	superSearch.Intelligence=     superDatabase.Intelligence 	 //intelligence
	superSearch.Power=            superDatabase.Power 	         //power
	superSearch.Occupation=       superDatabase.Occupation       //occupation
	superSearch.Image=            superDatabase.Image            //image
	superSearch.GroupAffiliation= superDatabase.GroupAffiliation //lista de grupos em que tal super está associado
	superSearch.Relatives=        len(superDatabase.Relatives)   //número de parentes
}