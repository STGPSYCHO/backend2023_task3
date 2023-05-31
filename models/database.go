package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=gorm password=gorm dbname=Go_Backend port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	// if db.Migrator().HasTable(&Category{}) {
	// 	db.Migrator().DropTable(&Category{})
	// }

	db.AutoMigrate(&User{}, &Blog{}, &Comment{}, &Category{}, &Tag{}, &Role{})

	DB = db
}
