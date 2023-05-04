package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"quackend/internal/model"
)

func main() {
	post := model.Post{}
	db, err := connectDB()
	db.First(&post)
	if err == nil {
		log.Printf("first post", post)
	}
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "12345678", "127.0.0.1:3306", "blog")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		log.Println("database connected")
		return db, err
	}
	log.Println("connect database failed")
	return nil, err
}
