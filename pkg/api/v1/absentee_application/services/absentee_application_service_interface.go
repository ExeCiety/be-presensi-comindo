package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/responses"

	"github.com/gofiber/fiber/v2"
)

type AbsenteeApplicationServiceInterface interface {
	FindAbsenteeApplications(c *fiber.Ctx, request *requests.GetAbsenteeApplications, responseData *[]responses.GetAbsenteeApplications) error
	FindAbsenteeApplication(c *fiber.Ctx, request *requests.GetAbsenteeApplication, responseData *responses.GetAbsenteeApplication) error
	CreateAbsenteeApplication(c *fiber.Ctx, request *requests.CreateAbsenteeApplication, responseData *responses.CreateAbsenteeApplication) error
	UpdateAbsenteeApplication(c *fiber.Ctx, request *requests.UpdateAbsenteeApplication, responseData *responses.UpdateAbsenteeApplication) error
	DeleteAbsenteeApplications(c *fiber.Ctx, request *requests.DeleteAbsenteeApplications, responseData *[]responses.DeleteAbsenteeApplications) error
}
