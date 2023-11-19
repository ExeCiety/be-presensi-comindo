package utils

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func SendApiResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, errors interface{}) error {
	return c.Status(statusCode).
		JSON(ApiResponse{
			Message: message,
			Data:    data,
			Errors:  errors,
		})
}
