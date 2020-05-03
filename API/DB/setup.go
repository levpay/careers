package DB

import (
	"API/models"
	"github.com/jinzhu/gorm"
	"os"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" port="+os.Getenv("DB_PORT")+
		"user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME")+"password="+os.Getenv("DB_PASSWORD"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(models.SuperOrVilan{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}
