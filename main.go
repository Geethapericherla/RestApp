package main

import (
	"fmt"

	"work/book"
	"work/database"
	"work/router"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDatabase() {
	var err error
	DBConn, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("DB not connected")
	}
	DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Connection successful")

}

func main() {

	app := fiber.New()
	InitDatabase()
	router.Setroutes(app)
	app.Listen(":8080")
	defer database.DBConn.Close()
}
