package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("createtodo", func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		todos = append(todos, todo)
		return c.JSON(&todo)

	})

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		tid := c.Params("id")
		var todo Todo
		for _, item := range todos {
			if item.Id == tid {
				// todo.Id = item.Id
				// todo.Id = item.Name
				// todo.Completed = item.Completed
				todo = *item
			} else {
				return c.SendString("connot found")
			}
		}
		return c.JSON(&todo)
	})
	log.Fatal(app.Listen(":3100"))

}
