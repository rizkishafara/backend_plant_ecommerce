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

	app.Get("/:jenis/:file", controller.GetFile)
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

func Product(app fiber.Router) {
	// app.Get("/product/getproduct", controller.GetProduct)
	// app.Get("/product/getproductbyid", controller.GetProductByID)
	// app.Get("/product/getproductbycategory", controller.GetProductByCategory)
	// app.Get("/product/getproductbysearch", controller.GetProductBySearch)
	// app.Get("/product/getproductbycategorysearch", controller.GetProductByCategorySearch)
	app.Post("/product/addproduct", controller.AddProduct)
	// app.Post("/product/updateproduct", controller.UpdateProduct)
	// app.Post("/product/deleteproduct", controller.DeleteProduct)
	// app.Post("/product/addproductimage", controller.AddProductImage)
	// app.Post("/product/updateproductimage", controller.UpdateProductImage)
	// app.Post("/product/deleteproductimage", controller.DeleteProductImage)
}
