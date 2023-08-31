package configs

import (
	"km5go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDatabase() {
	var dsn = "root:123ABC4d.@tcp(localhost:3306)/prakerja9"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn),  &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&models.News{})
}