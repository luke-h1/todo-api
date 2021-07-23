package todo

import (
	"github.com/gofiber/fiber"
	"github.com/luke-h1/todo-api/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func GetTodos(c *fiber.Ctx) {
	db := database.DBConn

	var todos []Todo
	db.Find(&todos)
	c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DBConn

	var todo Todo
	db.Find(&todo, id)
	c.JSON(todo)

}

func NewTodo(c *fiber.Ctx) {
	db := database.DBConn

	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&todo)
	c.JSON(todo)

}

func DeleteTodo(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DBConn

	var todo Todo

	db.First(&todo, id)
	if todo.Title == "" {
		c.Status(404).Send("No todo found with given id")
		return
	}
	db.Delete(&todo)
	c.Send("todo succesfully deleted")
}
