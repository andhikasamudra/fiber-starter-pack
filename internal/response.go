package internal

import "github.com/gofiber/fiber/v2"

func BuildResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
