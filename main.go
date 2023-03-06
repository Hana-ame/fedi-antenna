package main

import (
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"

	"github.com/Hana-ame/fedi-antenna/user"
	"github.com/Hana-ame/fedi-antenna/webfinger"
)

func main() {

	log.Println("Starting")

	app := fiber.New()

	// app.All(webfinger.WebFingerPath, func(c *fiber.Ctx) error {
	// 	return nil
	// })
	app.Mount(webfinger.WebFingerPath, webfingerApp())
	app.Use("/", func(c *fiber.Ctx) (err error) {
		log.Println(c.Path())
		// c.SendString(c.Path())
		return c.Next()
	})

	app.All("/users/:username", func(c *fiber.Ctx) (err error) {
		username := c.Params("username")
		if username == "" {
			return nil
		}

		as, err := user.UserAS(username)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Set("Content-Type", `application/ld+json; profile="https://www.w3.org/ns/activitystreams"; charset=utf-8`)
		_, err = c.Status(fiber.StatusOK).Write(as)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return nil
	})

	log.Fatal(app.Listen(":5000"))
}

func webfingerApp() *fiber.App {
	app := fiber.New()

	app.All("", func(c *fiber.Ctx) error {
		acct := c.Query("resource")

		acct, err := url.PathUnescape(acct)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "invalid_resource",
			})
		}

		// log.Println(acct)

		username, domain := webfinger.ParseAcct(acct)

		if !webfinger.CheckDomain(domain) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "invalid_resource",
			})
		}

		if !webfinger.CheckUserExist(username) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "no_such_user",
			})
		}

		err = c.Status(fiber.StatusOK).JSON(webfinger.GetResource(username, domain))

		return err
	})

	return app
}
