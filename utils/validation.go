package utils

import (
	"strings"

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

var MyValidation *validator.Validate

func BodyParserAndValidate(c *fiber.Ctx, request any) error {
	err := c.BodyParser(request)
	validationErrorMessages := GetValidationResult(ValidateStruct(request))

	if err != nil || len(validationErrorMessages) > 0 {
		return NewApiError(
			fiber.StatusUnprocessableEntity,
			Translate("err.validation_error", nil),
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
	var messageID string

	for _, v := range validationErrors {
		switch v.Tag {
		case "max", "min":
			messageID = v.Tag + "." + v.Field
			break
		default:
			messageID = v.Tag
			break
		}

		errResult[v.Field] = Translate("validation."+messageID, map[string]interface{}{
			"Field": v.Field,
			"Param": v.Param,
		})
	}

	return errResult
}
