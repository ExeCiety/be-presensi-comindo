package validations

import (
	"strings"

	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Namespace string
	Field     string
	Tag       string
	Param     string
	Value     interface{}
}

func (v *ValidationError) GetMessageId() string {
	messageId := v.Tag

	switch v.Tag {
	case enums.ValidationTagNameMax, enums.ValidationTagNameMin:
		messageId = v.Tag + "." + v.Field
		break
	}

	return messageId
}

func (v *ValidationError) GetParam() interface{} {
	param := v.Param

	switch v.Tag {
	case enums.ValidationTagNameDateGreaterThanField,
		enums.ValidationTagNameDateSameAsField:
		return GetDateGreaterThanFieldParam(v)
	}

	return param
}

var MyValidation *validator.Validate

func QueryParserAndValidate(c *fiber.Ctx, request any) error {
	err := c.QueryParser(request)
	validationErrorMessages := GetValidationResult(ValidateStruct(request))

	if err != nil || len(validationErrorMessages) > 0 {
		return utils.NewApiError(
			fiber.StatusUnprocessableEntity,
			utils.Translate("err.validation_error", nil),
			validationErrorMessages,
		)
	}

	return nil
}

func BodyParserAndValidate(c *fiber.Ctx, request any) error {
	err := c.BodyParser(request)
	validationErrorMessages := GetValidationResult(ValidateStruct(request))

	if err != nil || len(validationErrorMessages) > 0 {
		return utils.NewApiError(
			fiber.StatusUnprocessableEntity,
			utils.Translate("err.validation_error", nil),
			validationErrorMessages,
		)
	}

	return nil
}

func ValidateStruct(obj any) []*ValidationError {
	var validationErrors []*ValidationError

	if errValidate := MyValidation.Struct(obj); errValidate != nil {
		for _, validationError := range errValidate.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, &ValidationError{
				Namespace: strings.ToLower(validationError.Namespace()),
				Field:     strings.ToLower(validationError.Field()),
				Tag:       validationError.Tag(),
				Param:     validationError.Param(),
				Value:     validationError.Value(),
			})
		}
	}

	return validationErrors
}

func GetValidationResult(validationErrors []*ValidationError) map[string]string {
	errResult := make(map[string]string)
	for _, v := range validationErrors {
		errResult[v.Field] = utils.Translate("validation."+v.GetMessageId(), map[string]interface{}{
			"Field": v.Field,
			"Param": v.GetParam(),
			"Value": v.Value,
		})
	}

	return errResult
}
