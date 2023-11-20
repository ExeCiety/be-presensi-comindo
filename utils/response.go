package utils

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type PaginationResponse struct {
	Meta MetaPagination `json:"meta"`
	Data interface{}    `json:"data"`
}

type MetaPagination struct {
	TotalData   int64 `json:"total_data"`
	TotalPage   int64 `json:"total_page"`
	PageSize    int   `json:"page_size"`
	CurrentPage int   `json:"current_page"`
}

func SendApiResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, errors interface{}) error {
	return c.Status(statusCode).
		JSON(ApiResponse{
			Message: message,
			Data:    data,
			Errors:  errors,
		})
}

func GetResourceResponseData(data interface{}, request PaginationRequest) interface{} {
	if request.Paginate == true {
		return PaginationResponse{
			Meta: MetaPagination{
				TotalData:   request.TotalData,
				TotalPage:   int64(math.Ceil(float64(request.TotalData) / float64(request.PerPage))),
				PageSize:    request.PerPage,
				CurrentPage: request.Page,
			},
			Data: data,
		}
	}

	return data
}
