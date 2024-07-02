package main

import (
	"fmt"
	"log"
	"tanaman/config"
	"tanaman/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	config := config.LoadConfig(".")

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Set CORS headers

		

		log.Fatalf("Origin: %s", c.Get("Origin"))

		origin := c.Get("Origin")

		if origin == "http://localhost:3000" || origin == "https://planting-ecommerce.vercel.app" {
			c.Set("Access-Control-Allow-Origin", origin)
		}
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		// Go to next middleware:
		return c.Next()
	})
	signingKey := []byte("@9r33n3l394nt123")
	configjwt := jwtware.Config{
		TokenLookup:  "header:Authorization",
		ErrorHandler: app.ErrorHandler,
		SigningKey:   signingKey,
	}

	app2 := app.Group("/v1")
	app2.Use(jwtware.New(configjwt))

	routes.Home(app)
	routes.Catalog(app)
	routes.Auth(app)
	routes.Reference(app)
	routes.Profile(app2)
	routes.Cart(app2)
	routes.Admin(app2)

	host := fmt.Sprintf(":%d", config.ServerPort)
	log.Fatal(app.Listen(host))
}
