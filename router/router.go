package router

import (
	"work/book"

	"github.com/gofiber/fiber/v2"
)

func Setroutes(app *fiber.App) {

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.CreateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}
