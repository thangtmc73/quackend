package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"quackend/internal/model"
)

func main() {
	post := model.Post{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "12345678", "127.0.0.1:3306", "blog")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.First(&post)
	if err == nil {
		fmt.Println("database connected")
		fmt.Printf("first post", post)
	} else {
		fmt.Println("connect database failed")
	}
}
