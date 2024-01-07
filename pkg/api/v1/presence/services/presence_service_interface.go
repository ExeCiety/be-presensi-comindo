package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/responses"

	"github.com/gofiber/fiber/v2"
)

type PresenceServiceInterface interface {
	FindPresences(c *fiber.Ctx, request *requests.GetPresences, responseData *[]responses.GetPresences) error
	FindPresence(c *fiber.Ctx, request *requests.GetPresence, responseData *responses.GetPresence) error
	CreatePresence(c *fiber.Ctx, request *requests.CreatePresence, responseData *responses.CreatePresence) error
	UpdatePresence(c *fiber.Ctx, request *requests.UpdatePresence, responseData *responses.UpdatePresence) error
	DeletePresences(c *fiber.Ctx, request *requests.DeletePresences, responseData *[]responses.DeletePresences) error
}
