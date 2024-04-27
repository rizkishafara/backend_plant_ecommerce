package controller

import (
	"fmt"
	"regexp"
	"tanaman/model"
	"tanaman/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
}
func Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullname := c.FormValue("fullname")

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		response.Success = false
		response.Message = "Email tidak valid"
		return c.JSON(response)
	}

	if email == "" || password == "" || fullname == "" {
		response.Success = false
		response.Message = "Data wajib diisi"
		return c.JSON(response)
	}

	createUUid := uuid.New()
	data := model.Register(email, password, fullname, time.Now().Format("2006-01-02"), createUUid.String())

	return c.JSON(data)
}
