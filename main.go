package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/luke-h1/todo-api/database"
	"github.com/luke-h1/todo-api/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/todo", todo.GetTodos)
	app.Get("/api/todo/:id", todo.GetTodo)
	app.Post("/api/todo", todo.NewTodo)
	app.Delete("/api/todo/:id", todo.DeleteTodo)

}

func createConn() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		panic("❌ Failed to connect to DB!")
	}
	fmt.Println("✅ Succesfully connected to DB")
}

func main() {
	app := fiber.New()

	createConn()

	// close DB conns as soon as possible

	setupRoutes(app)

	app.Listen(4000)
}
