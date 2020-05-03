package DB

import (
	"API/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+
		" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME") +" sslmode=disable"+" password="+os.Getenv("DB_PASSWORD")+"")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(models.SuperOrVilan{})

	return db
}
