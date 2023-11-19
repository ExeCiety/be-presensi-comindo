package utils

import (
	"github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

func (ae ApiError) Error() string {
	return ae.Message
}

func NewApiError(code int, message string, errors any) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case *ApiError:
		statusMessage := err.(*ApiError).Error()
		if err.(*ApiError).Code == 0 || err.(*ApiError).Code == fiber.StatusInternalServerError {
			statusMessage = enums.StatusMessageInternalServerError
		}

		return SendApiResponse(c, err.(*ApiError).Code, statusMessage, nil, err.(*ApiError).Errors)
	default:
		return c.Status(fiber.StatusInternalServerError).SendString(enums.StatusMessageInternalServerError)
	}
}
