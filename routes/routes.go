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
	app.Get("/profile/uservoucher", controller.GetUserVoucher)

	app.Post("/profile/addaddress", controller.AddAddress)
	app.Post("/profile/updateaddress", controller.UpdateAddress)
	app.Post("/profile/deleteaddress", controller.DeleteAddress)
	app.Get("/profile/getaddress", controller.GetAddress)
}

func Catalog(app *fiber.App) {

	app.Get("/product", controller.GetProduct)
	app.Get("/product/category", controller.GetProductCategory)
	app.Get("/product/size", controller.GetProductSize)
	app.Get("/product/detail", controller.GetProductDetail)
	app.Get("/product/review", controller.GetProductReview)
	// app.Get("/product/getproductbyid", controller.GetProductByID)
	// app.Get("/product/getproductbycategory", controller.GetProductByCategory)
	// app.Get("/product/getproductbysearch", controller.GetProductBySearch)
	// app.Get("/product/getproductbycategorysearch", controller.GetProductByCategorySearch)

}

func Cart(app fiber.Router) {
	app.Post("/cart/addtocart", controller.AddToCart)
	app.Post("/cart/updatecart", controller.UpdateCart)
	app.Post("/cart/deletecart", controller.DeleteCart)
	app.Get("/cart/getcart", controller.GetCart)
	app.Get("/cart/getcountcart", controller.GetCountCart)
}
func Reference(app *fiber.App) {
	app.Get("/reference/province", controller.GetProvince)
	app.Get("/reference/city", controller.GetCity)
	app.Get("/reference/district", controller.GetDistrict)
	app.Get("/reference/village", controller.GetVillage)
	app.Get("/reference/postalcode", controller.GetPostalCode)
	app.Get("/reference/shipping", controller.GetShipping)
	app.Get("/reference/payment", controller.GetPayment)
}

// admin route
func Admin(app fiber.Router) {
	// produk
	app.Post("/product/addproduct", controlleradmin.AddProduct)
	app.Post("/product/updateproduct", controlleradmin.UpdateProduct)
	app.Post("/product/deleteproduct", controlleradmin.DeleteProduct)
	// app.Post("/product/addproductimage", controller.AddProductImage)
	// app.Post("/product/updateproductimage", controller.UpdateProductImage)
	// app.Post("/product/deleteproductimage", controller.DeleteProductImage)

	// referensi
	app.Post("/reference/addcategoryproduct", controlleradmin.AddCategoryProduct)
}
