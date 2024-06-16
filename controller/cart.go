package controller

import (
	"tanaman/model"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func AddToCart(c *fiber.Ctx) error {

	product_id := c.FormValue("product_id")
	qty := c.FormValue("qty")
	size_id := c.FormValue("size_id")
	user_id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	if product_id == "" || qty == "" {
		response.Status = 404
		response.Message = "Product ID and Qty must be filled"
		return c.JSON(response)
	}

	return c.JSON(model.AddToCart(uuid.New().String(), product_id, qty, size_id, user_id, time.Now().Format("2006-01-02")))
}
func UpdateCart(c *fiber.Ctx) error {
	id := c.FormValue("cart_id")
	qty := c.FormValue("qty")
	size_id := c.FormValue("size_id")

	if id == "" {
		response.Status = 404
		response.Message = "Cart ID must be filled"
		return c.JSON(response)
	}

	return c.JSON(model.UpdateCart(id, qty, size_id, time.Now().Format("2006-01-02")))
}
func DeleteCart(c *fiber.Ctx) error {
	id := c.FormValue("cart_id")

	if id == "" {
		response.Status = 404
		response.Message = "Cart ID must be filled"
		return c.JSON(response)
	}

	return c.JSON(model.DeleteCart(id))
}