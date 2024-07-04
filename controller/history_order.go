package controller

import (
	"tanaman/model"
	"tanaman/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetHistoryOrder(c *fiber.Ctx) error {

	user_id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")

	return c.JSON(model.GetListHistoryOrder(user_id))
}
