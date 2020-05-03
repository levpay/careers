package controllers

import (
	"API/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
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
