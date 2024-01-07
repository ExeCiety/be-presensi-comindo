package controllers

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/responses"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/services"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

type PresenceController struct {
	PresenceService services.PresenceServiceInterface
}

func NewPresenceController(presenceService services.PresenceServiceInterface) *PresenceController {
	return &PresenceController{
		PresenceService: presenceService,
	}
}

func (pc *PresenceController) FindPresences(c *fiber.Ctx) error {
	request := new(requests.GetPresences)
	responseData := new([]responses.GetPresences)

	if err := pc.PresenceService.FindPresences(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("presence.get_presences_success", nil),
		utils.GetResourceResponseData(responseData, request.PaginationRequest),
		nil,
	)
}

func (pc *PresenceController) FindPresence(c *fiber.Ctx) error {
	request := new(requests.GetPresence)
	responseData := new(responses.GetPresence)

	if err := pc.PresenceService.FindPresence(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("presence.get_presences_success", nil),
		responseData,
		nil,
	)
}

func (pc *PresenceController) CreatePresence(c *fiber.Ctx) error {
	request := new(requests.CreatePresence)
	responseData := new(responses.CreatePresence)

	if err := pc.PresenceService.CreatePresence(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("presence.create_presence_success", nil),
		responseData,
		nil,
	)
}

func (pc *PresenceController) UpdatePresence(c *fiber.Ctx) error {
	request := new(requests.UpdatePresence)
	responseData := new(responses.UpdatePresence)

	if err := pc.PresenceService.UpdatePresence(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("presence.update_presence_success", nil),
		responseData,
		nil,
	)
}

func (pc *PresenceController) DeletePresences(c *fiber.Ctx) error {
	request := new(requests.DeletePresences)
	responseData := new([]responses.DeletePresences)

	if err := pc.PresenceService.DeletePresences(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("presence.delete_presences_success", nil),
		responseData,
		nil,
	)
}
