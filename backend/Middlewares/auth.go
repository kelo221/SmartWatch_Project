package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

const SecretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
}

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "not logged in",
		})
	}

	return c.Next()
}

func GetUserID(c *fiber.Ctx) (string, error) {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "", err

	}

	payload := token.Claims.(*ClaimsWithScope)

	return payload.Subject, nil

}

func GenerateJWT(id string) (string, error) {

	var payload ClaimsWithScope

	payload.Subject = id
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))

}
