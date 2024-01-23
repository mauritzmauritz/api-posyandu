package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itsLeonB/posyandu-api/core/exception"
)

func AuthorizeUser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			panic(exception.BadRequestError{
				Message: "Invalid parameter",
			})
		}

		claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		user := int(claims["id"].(float64))

		if role == "admin" {
			return c.Next()
		}

		if user == id {
			return c.Next()
		} else {
			panic(exception.UnauthorizedError{
				Message: "Unauthorized access!",
			})
		}
	}
}

func AuthorizeRole() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role := c.Params("role")

		claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
		user := claims["role"].(string)

		if user == "admin" || user == "bidan" {
			return c.Next()
		} else {
			if role != "admin" {
				return c.Next()
			} else {
				panic(exception.UnauthorizedError{
					Message: "Unauthorized access!",
				})
			}
		}
	}
}
