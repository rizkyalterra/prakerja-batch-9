package configs

import (
	"fmt"
	"km5go/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDatabase() {
	
	var dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASSWORD"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
		)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn),  &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&models.News{}, &models.User{})
}