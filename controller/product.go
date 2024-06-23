package controller

import (
	"encoding/json"
	"fmt"
	"tanaman/model"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {

	sizeData := c.Query("length")
	page := c.Query("page")
	search := c.Query("search")
	minPrice := c.Query("min-price")
	maxPrice := c.Query("max-price")
	sort := c.Query("sort")

	idCategory := c.FormValue("category")
	idSize := c.FormValue("size")

	if page == "" {
		page = "1"
	}

	var arrayCategory []int

	if idCategory != "" {
		err := json.Unmarshal([]byte(idCategory), &arrayCategory)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			response.Status = 500
			response.Message = "Error unmarshalling JSON Category:" + err.Error()
			return c.JSON(response)
		}
	}

	product := model.GetProduct(sizeData, page, search, minPrice, maxPrice, sort, arrayCategory, idSize)
	return c.JSON(product)
}
func GetProductCategory(c *fiber.Ctx) error {
	return c.JSON(model.GetProductCategory())
}
func GetProductSize(c *fiber.Ctx) error {
	return c.JSON(model.GetProductSize())
}
func GetProductDetail(c *fiber.Ctx) error {
	uuid := c.Query("id")
	if uuid == "" {
		response.Status = 404
		response.Message = "UUID is required"
		return c.JSON(response)
	}
	return c.JSON(model.GetProductDetail(uuid))
}
func GetProductReview(c *fiber.Ctx) error {
	uuid := c.Query("id")
	if uuid == "" {
		response.Status = 404
		response.Message = "Product ID is required"
		return c.JSON(response)
	}
	return c.JSON(model.GetProductReview(uuid))
}
