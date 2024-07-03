package controller

// import (
// 	"fmt"
// 	"math/rand"
// 	"tanaman/model"
// 	"tanaman/utils"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/google/uuid"
// )

// func PostCheckout(c *fiber.Ctx) error {
// 	user_id := utils.GetValJWT(c.Locals("user").(*jwt.Token), "idreq")
// 	order_number := generateOrderNumber()
// 	payment_id := c.FormValue("payment_id")
// 	// json_alamat := c.FormValue("json_alamat")
// 	notes := c.FormValue("notes")
// 	resi := c.FormValue("resi")
// 	sub_total := c.FormValue("sub_total")
// 	discount := c.FormValue("discount")
// 	idshipping := c.FormValue("idshipping")

// 	var alamat []string
// 	chekout := model.PostCheckout(uuid.New().String(), user_id.(string), "1", order_number, payment_id, notes, resi, sub_total, discount, idshipping, "0", time.Now().Format("2006-01-02"), alamat)
// 	return c.JSON(chekout)
// }

// func generateOrderNumber() string {
// 	// Get current date and time
// 	currentTime := time.Now()
// 	dateTimeFormat := currentTime.Format("02012006150405") // ddmmyyyyhhmmss

// 	// Generate a unique identifier (xxx part)
// 	rand.Seed(time.Now().UnixNano())
// 	uniqueID := rand.Intn(1000) // random number between 0 and 999

// 	// Combine them into the desired format
// 	orderNumber := fmt.Sprintf("INV/%s%03d", dateTimeFormat, uniqueID)

// 	return orderNumber
// }
