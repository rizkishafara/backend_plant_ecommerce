package controller

import (
	"fmt"
	"math/rand"
	"time"

	"tanaman/model"

	"github.com/gofiber/fiber/v2"
)

func PostCheckout(c *fiber.Ctx) error {

	order_number := generateOrderNumber()
	payment_id := c.FormValue("payment_id")
	json_alamat := c.FormValue("json_alamat")
	notes := c.FormValue("notes")
	resi := c.FormValue("resi")
	sub_total := c.FormValue("sub_total")
	discount := c.FormValue("discount")
	idshipping := c.FormValue("idshipping")

	var alamat []string
	chekout := model.Checkout(order_number, payment_id, json_alamat, notes, resi, sub_total, discount, idshipping, time.Now().Format("2006-01-02"),"","","",alamat)

	return c.JSON(chekout)
}

func generateOrderNumber() string {
	// Get current date and time
	currentTime := time.Now()
	dateTimeFormat := currentTime.Format("02012006150405") // ddmmyyyyhhmmss

	// Generate a unique identifier (xxx part)
	rand.Seed(time.Now().UnixNano())
	uniqueID := rand.Intn(1000) // random number between 0 and 999

	// Combine them into the desired format
	orderNumber := fmt.Sprintf("INV/%s%03d", dateTimeFormat, uniqueID)

	return orderNumber
}
