package controller

import (
	"tanaman/model"

	"github.com/gofiber/fiber/v2"
)

func GetBanner(c *fiber.Ctx) error {
	banner := model.GetBanner()
	return c.JSON(banner)
	// return c.JSON("banner")
}
func GetPopularCategory(c *fiber.Ctx) error {
	category := model.GetPopularCategory()
	return c.JSON(category)
}
func GetNewestProduct(c *fiber.Ctx) error {
	product := model.GetNewestProduct()
	return c.JSON(product)
}
