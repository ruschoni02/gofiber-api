package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/", hello)
	app.Get("/accounts/:accountId/:symbol-:bla/orders", bla)
	app.Get("/accounts/:accountId/:symbol/orders", orders)
	app.Get("/accounts/:accountId/:symbol", symbol)
	app.Get("/accounts/:accountId", accounts)

	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{Message: "Hello, World"})
}

func accounts(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{
		Message: fmt.Sprintf("Account: %s", c.Params("accountId")),
	})
}

func symbol(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{
		Message: fmt.Sprintf(
			"Account: %s | Symbol: %s",
			c.Params("accountId"),
			c.Params("symbol"),
		),
	})
}

func orders(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{
		Message: fmt.Sprintf(
			"Account: %s | Symbol: %s | Orders",
			c.Params("accountId"),
			c.Params("symbol"),
		),
	})
}

func bla(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{
		Message: fmt.Sprintf(
			"Account: %s | Symbol: %s | Bla: %s | Orders",
			c.Params("accountId"),
			c.Params("symbol"),
			c.Params("bla"),
		),
	})
}
