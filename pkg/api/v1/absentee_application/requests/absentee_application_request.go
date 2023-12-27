package requests

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type GetAbsenteeApplications struct {
	UserId    *string `query:"user_id" validate:"omitempty,uuid,exists=users;id"`
	DateStart *string `query:"date_start" validate:"omitempty,datetime=2006-01-02"`
	DateEnd   *string `query:"date_end" validate:"omitempty,datetime=2006-01-02"`
	utils.PaginationRequest
}

type GetAbsenteeApplication struct {
	Id     string
	UserId *string `query:"-"`
}

type CheckIfAbsenteeApplicationExistOnThatDays struct {
	ExceptionId *uuid.UUID
	UserId      uuid.UUID
	DateStart   string
	DateEnd     string
}

type CreateAbsenteeApplication struct {
	UserId     string `json:"user_id" validate:"required,uuid,exists=users;id"`
	Type       string `json:"type" validate:"required,oneof=sick permission paid_leave"`
	Status     string `json:"status" validate:"omitempty,oneof=approved rejected in_review"`
	DateStart  string `json:"date_start" validate:"required,datetime=2006-01-02,date_greater_than_today=2006-01-02"`
	DateEnd    string `json:"date_end" validate:"required,datetime=2006-01-02,date_greater_than_field=date start;2006-01-02"`
	Reason     string `json:"reason" validate:"required_if=Type permission,required_if=Type paid_leave"`
	Attachment string `json:"attachment" validate:"required_if=Type sick"`
}

type UpdateAbsenteeApplication struct {
	Id         string `json:"-"`
	UserId     string `json:"user_id" validate:"omitempty,uuid,exists=users;id"`
	Type       string `json:"type" validate:"omitempty,oneof=sick permission paid_leave"`
	DateStart  string `json:"date_start" validate:"required_with=DateEnd,omitempty,datetime=2006-01-02,date_greater_than_today=2006-01-02"`
	DateEnd    string `json:"date_end" validate:"required_with=DateStart,omitempty,datetime=2006-01-02,date_greater_than_field=date start;2006-01-02"`
	Status     string `json:"status" validate:"omitempty,oneof=approved rejected in_review"`
	Reason     string `json:"reason"`
	Attachment string `json:"attachment"`
}

type DeleteAbsenteeApplications struct {
	Ids    []string `json:"ids" validate:"required,min=1,dive,uuid,exists=absentee_applications;id"`
	UserId string   `json:"-"`
	Status string   `json:"-"`
}
