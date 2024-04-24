package book

import (
	"work/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	books := []Book{}
	d := database.DBConn
	d.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	var book Book
	p := c.Params("id")
	d := database.DBConn
	d.Find(&book, p)
	if book.Name == "" {
		return c.Status(501).Send([]byte("book not found"))
	}
	return c.JSON(book)
}
func CreateBook(c *fiber.Ctx) error {
	var book Book
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(503).Send([]byte("body parsing failed"))
	}
	d := database.DBConn
	d.Create(&book)
	return c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) error {
	var book Book
	p := c.Params("id")
	d := database.DBConn
	d.First(&book, p)
	if book.Name == "" {
		return c.Status(501).Send([]byte("book not found"))
	}
	return c.Send([]byte("Book deleted"))
}
