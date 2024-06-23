package controller

import (
	"tanaman/model"

	"github.com/gofiber/fiber/v2"
)

func GetProvince(c *fiber.Ctx) error {

	return c.JSON(model.GetProvince())
}
func GetCity(c *fiber.Ctx) error {
	province_id := c.Query("province_id")
	return c.JSON(model.GetCity(province_id))
}
func GetDistrict(c *fiber.Ctx) error {
	city_id := c.Query("city_id")
	return c.JSON(model.GetDistrict(city_id))
}
func GetVillage(c *fiber.Ctx) error {
	district_id := c.Query("district_id")
	return c.JSON(model.GetVillage(district_id))
}
func GetPostalCode(c *fiber.Ctx) error {
	village_id := c.Query("village_id")
	return c.JSON(model.GetPostalCode(village_id))
}
func GetShipping(c *fiber.Ctx) error {
	return c.JSON(model.GetShipping())
}
