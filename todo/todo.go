package todo

import (
	"github.com/gofiber/fiber"
)

func GetTodos(c *fiber.Ctx) {
	c.Send("All todos")
}

func GetTodo(c *fiber.Ctx) {
	c.Send("single todo")
}

func NewTodo(c *fiber.Ctx) {
	c.Send("new todo")
}

func DeleteTodo(c *fiber.Ctx) {
	c.Send("delete todo")
}
