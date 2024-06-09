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
	var arraySize []int

	if idCategory != "" {
		err := json.Unmarshal([]byte(idCategory), &arrayCategory)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			response.Status = 500
			response.Message = "Error unmarshalling JSON Category:" + err.Error()
			return c.JSON(response)
		}
	}
	if idSize != "" {
		err := json.Unmarshal([]byte(idSize), &arraySize)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			response.Status = 500
			response.Message = "Error unmarshalling JSON Size:" + err.Error()
			return c.JSON(response)
		}
	}

	product := model.GetProduct(sizeData, page, search, minPrice, maxPrice, sort, arrayCategory, arraySize)
	return c.JSON(product)
}
