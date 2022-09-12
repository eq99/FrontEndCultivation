package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {
	// Initialize standard Go html template engine
	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", GetHome)
	app.Get("/signup", GetSignup)
	app.Post("/signup", PostSignup)

	log.Fatal(app.Listen(":3000"))
}

func GetHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "欢迎来到 Bilibili 大学",
	})
}

func GetSignup(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"Title": "Hello, World!",
	})
}

func PostSignup(c *fiber.Ctx) error {
	password := c.FormValue("password")

	return c.Render("profile", fiber.Map{
		"Title":    "我的主页",
		"Password": password,
	})

}
