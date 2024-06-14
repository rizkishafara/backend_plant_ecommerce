package routes

import (
	"tanaman/controller"
	controlleradmin "tanaman/controller_admin"

	"github.com/gofiber/fiber/v2"
)
func Home(app *fiber.App) {
	app.Get("/home/banner", controller.GetBanner)
	app.Get("/home/popular-category", controller.GetPopularCategory)
	app.Get("/home/newest-product", controller.GetNewestProduct)
}

func Auth(app *fiber.App) {
	app.Post("/auth/login", controller.Login)
	app.Post("/auth/register", controller.Register)
	app.Post("/auth/forgotpassword", controller.ForgotPassword)
	app.Post("/auth/updatepassword", controller.UpdatePassword)

	app.Get("/file/:jenis/:file", controller.GetFile)
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

func Product(app *fiber.App) {

	app.Get("/product", controller.GetProduct)
	// app.Get("/product/getproductbyid", controller.GetProductByID)
	// app.Get("/product/getproductbycategory", controller.GetProductByCategory)
	// app.Get("/product/getproductbysearch", controller.GetProductBySearch)
	// app.Get("/product/getproductbycategorysearch", controller.GetProductByCategorySearch)

	// FOR ADMIN ONLY
	app.Post("/product/addproduct", controlleradmin.AddProduct)
	app.Post("/product/updateproduct", controlleradmin.UpdateProduct)
	app.Post("/product/deleteproduct", controlleradmin.DeleteProduct)
	// app.Post("/product/addproductimage", controller.AddProductImage)
	// app.Post("/product/updateproductimage", controller.UpdateProductImage)
	// app.Post("/product/deleteproductimage", controller.DeleteProductImage)
}
