package authControllers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	db "github.com/miceremwirigi/go-sales-api/config/dbConnection"
	"github.com/miceremwirigi/go-sales-api/models"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "invalid post request",
			},
		)
	}

	// check postcode field is empty or not
	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is required",
				"data":    data,
			},
		)
	}
	var cashier models.Cashier
	db.DB.Where("Id = ?", cashierId).First(&cashier)

	// validation
	if cashier.Id == 0 {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
				"error":   map[string]interface{}{},
			},
		)
	}

	if cashier.Passcode != data["passcode"] {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode not match",
				"error":   map[string]interface{}{},
			},
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), // 1 day
	},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Token expired or Invalid",
			},
		)
	}

	cashierData := map[string]interface{}{}
	cashierData["token"] = tokenString

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Login successful",
			"data":    cashierData,
		},
	)
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func Passcode(c *fiber.Ctx) error {
	return nil
}
