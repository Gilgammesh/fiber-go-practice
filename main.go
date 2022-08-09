package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

func getUsers(c *fiber.Ctx) error {
	user := User{
		Id:        uuid.NewString(),
		Firstname: "Carlos",
		Lastname:  "Santander",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())

	// Routes

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server ready")
	})

	// app.Get("/users", getUsers)
	// app.Post("/users", createUser)

	app.Use(requestid.New())

	// Group Routes
	userGroup := app.Group("/users")
	userGroup.Get("", getUsers)
	userGroup.Post("", createUser)

	app.Listen(":3000")
}
