package controllers

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/responses"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/services"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

type AbsenteeApplicationController struct {
	AbsenteeApplicationService services.AbsenteeApplicationServiceInterface
}

func NewAbsenteeApplicationController(
	absenteeApplicationService services.AbsenteeApplicationServiceInterface,
) *AbsenteeApplicationController {
	return &AbsenteeApplicationController{
		AbsenteeApplicationService: absenteeApplicationService,
	}
}

func (apc *AbsenteeApplicationController) FindAbsenteeApplications(c *fiber.Ctx) error {
	request := new(requests.GetAbsenteeApplications)
	responseData := new([]responses.GetAbsenteeApplications)

	if err := apc.AbsenteeApplicationService.FindAbsenteeApplications(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("absentee_application.get_absentee_applications_success", nil),
		utils.GetResourceResponseData(responseData, request.PaginationRequest),
		nil,
	)
}

func (apc *AbsenteeApplicationController) FindAbsenteeApplication(c *fiber.Ctx) error {
	request := new(requests.GetAbsenteeApplication)
	responseData := new(responses.GetAbsenteeApplication)

	if err := apc.AbsenteeApplicationService.FindAbsenteeApplication(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("absentee_application.get_absentee_applications_success", nil),
		responseData,
		nil,
	)
}

func (apc *AbsenteeApplicationController) CreateAbsenteeApplication(c *fiber.Ctx) error {
	request := new(requests.CreateAbsenteeApplication)
	responseData := new(responses.CreateAbsenteeApplication)

	if err := apc.AbsenteeApplicationService.CreateAbsenteeApplication(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("absentee_application.create_absentee_application_success", nil),
		responseData,
		nil,
	)
}

func (apc *AbsenteeApplicationController) UpdateAbsenteeApplication(c *fiber.Ctx) error {
	request := new(requests.UpdateAbsenteeApplication)
	responseData := new(responses.UpdateAbsenteeApplication)

	if err := apc.AbsenteeApplicationService.UpdateAbsenteeApplication(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("absentee_application.update_absentee_application_success", nil),
		responseData,
		nil,
	)
}

func (apc *AbsenteeApplicationController) DeleteAbsenteeApplications(c *fiber.Ctx) error {
	request := new(requests.DeleteAbsenteeApplications)
	responseData := new([]responses.DeleteAbsenteeApplications)

	if err := apc.AbsenteeApplicationService.DeleteAbsenteeApplications(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("absentee_application.delete_absentee_applications_success", nil),
		responseData,
		nil,
	)
}
