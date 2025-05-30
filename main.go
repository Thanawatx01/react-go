package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, World")

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loaing .env file")
	}

	PORT := "4000"

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// create todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is Required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(200).JSON(todo)
	})

	// update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todo not find"})
	})

	// delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "Success"})
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todo not find"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
