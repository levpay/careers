package controllers

import (
	"API/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"

	"time"
)

//Cadastrar um Super/Vilão
func CreateSuperOrVilan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.SuperOrVilanInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var superOrVilan models.SuperOrVilan
	//Cada super deve ser cadastrado somente a partir do seu name.
	if db.Where("name = ? AND deleted = ?", input.Name,"0").Find(&superOrVilan).RecordNotFound() {
			superOrVilan.Name=             input.Name
			superOrVilan.FullName=         input.FullName
			superOrVilan.Alignment=        input.Alignment
			superOrVilan.Intelligence=     input.Intelligence
			superOrVilan.Power=            input.Power
			superOrVilan.Occupation=       input.Occupation
			superOrVilan.Image=            input.Image
			superOrVilan.GroupAffiliation= input.GroupAffiliation
			superOrVilan.Relatives=        input.Relatives
		if err := db.Create(&superOrVilan).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{"success": "Super or Vilan criado !"})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": "registro existente",
	})
	return
}
//Listar todos os Super's cadastrados
func ListAllSuperOrVilan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var supers []models.SuperOrVilan
	if err := db.Find(&supers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": supers,
	})
	return
}
func ListAllSupers(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var supers []models.SuperOrVilan
	if err:=db.Where("alignment = ?","good").Find(&supers).Error;err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": supers,
	})
	return
}
func ListAllVilans(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var supers []models.SuperOrVilan
	if err:=db.Where("alignment = ?","bad").Find(&supers).Error;err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": supers,
	})
	return
}

//Buscar por nome
func SearchByName(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var superDatabase models.SuperOrVilan
	/*
	if err:=dataBaseSearch("name",c.Query("name"),c,&superDatabase);err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	*/
	fmt.Println(c.Query("name"))
	if err:=db.Where("name = ? AND deleted = ?",c.Query("name"),false).Find(&superDatabase).Error;err!=nil{
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"success":population(superDatabase)})
	return
}
//Buscar por 'uuid'
func SearchByUuid(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var superDatabase models.SuperOrVilan
	/*
	if err:=dataBaseSearch("name",c.Query("uuid"),c,&superDatabase);err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	*/
	if err:=db.Where("uuid = ? AND deleted = ?",c.Query("uuid"),false).Find(&superDatabase).Error;err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"success":population(superDatabase)})
	return
}
//Remover o Super
func RemoveSuper(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var superDatabase models.SuperOrVilan
	/*if err:=dataBaseSearch("name",c.Query("uuid"),c,&superDatabase);err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}*/
	if err:=db.Where("name = ? AND deleted = ?",c.Query("name"),false).Find(&superDatabase).Error;err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
	}

	if err :=db.Model(&superDatabase).Update(map[string]interface{}{
		"deleted": true,
		"deleted_at": time.Now().Unix(),
	}).Error;err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"success":"user soft removed by name"})
	return
}
//FieldOfSearch
/*func dataBaseSearch(fieldOfSearch string,Query string,c *gin.Context, superDatabase*models.SuperOrVilan) error{
	db := c.MustGet("db").(*gorm.DB)
	fieldOfSearch=strings.ToLower(fieldOfSearch)
	if fieldOfSearch != "name" || fieldOfSearch !="uuid"{
		return errors.New("parametro invalido")
	}
	//Recuperando os dados do BD
	if err:=db.Where(""+ fieldOfSearch+" = ? AND deleted = ?",Query,false).Find(superDatabase).Error;err!=nil{
		return  errors.New(err.Error())
	}
	return nil
}*/
//esta função popula e devolve para quem chamou
func population(superDatabase models.SuperOrVilan) models.SuperOrVilanSearch {
	var superSearch models.SuperOrVilanSearch
	superSearch.Uuid=             superDatabase.Uuid             //uuid
	superSearch.Name=             superDatabase.Name             //name
	superSearch.FullName=         superDatabase.FullName         //full name
	superSearch.Intelligence=     superDatabase.Intelligence 	 //intelligence
	superSearch.Power=            superDatabase.Power 	         //power
	superSearch.Occupation=       superDatabase.Occupation       //occupation
	superSearch.Image=            superDatabase.Image            //image
	superSearch.GroupAffiliation= superDatabase.GroupAffiliation //lista de grupos em que tal super está associado
	superSearch.Relatives=        len(strings.Split(superDatabase.Relatives,";"))   //número de parentes

	return superSearch
}