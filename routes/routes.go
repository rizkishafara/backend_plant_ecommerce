package routes

import (
	"tanaman/controller"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	app.Post("/v1/auth/login", controller.Login)
	app.Post("/v1/auth/register", controller.Register)
	// app.Post("/auth/login", controllers.Login)
	// app.Post("/auth/lupapin", controllers.Lupapin)
	// app.Post("/auth/updateprofil", controllers.Updateprofile)
}
