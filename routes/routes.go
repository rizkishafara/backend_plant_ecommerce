package routes

import (
	"tanaman/controller"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	app.Post("/auth/login", controller.Login)
	app.Post("/auth/register", controller.Register)
	app.Post("/auth/forgotpassword", controller.ForgotPassword)
	app.Post("/auth/updatepassword", controller.UpdatePassword)
	// app.Post("/auth/login", controllers.Login)
	// app.Post("/auth/lupapin", controllers.Lupapin)
	// app.Post("/auth/updateprofil", controllers.Updateprofile)
}
func Profile(app fiber.Router) {
	app.Get("/profile/getprofile", controller.GetProfileUser)
	app.Post("/profile/updateprofile", controller.UpdateProfileUser)
	app.Get("/profile/getcountvoucher", controller.GetCountVoucher)
	app.Get("/profile/getcountloyalty", controller.GetCountLoyalty)
}
