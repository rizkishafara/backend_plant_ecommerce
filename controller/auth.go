package controller

import (
	"fmt"
	"tanaman/model"
	"tanaman/utils"

	"github.com/gofiber/fiber/v2"
)

var response utils.Respon

func Login(c *fiber.Ctx) error {
	fmt.Println("Login")
	// return c.JSON(response)
	// BODY
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		response.Success = false
		response.Message = "Email dan Password wajib diisi"
		return c.JSON(response)
	}
	data := model.Login(email, password)

	return c.JSON(data)

	// return c.JSON(data)
}
