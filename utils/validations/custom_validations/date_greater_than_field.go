package custom_validations

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func DateGreaterThanField(fl validator.FieldLevel) bool {
	paramValues := strings.SplitN(fl.Param(), ";", 2)
	if len(paramValues) <= 0 {
		return false
	}

	fieldName := paramValues[0]
	if fieldName == "" {
		return false
	}

	layout := "2006-01-02"
	if len(paramValues) > 1 {
		layout = paramValues[1]
	}

	parse, err := time.Parse(layout, fl.Field().String())
	if err != nil {
		log.Error(err)
		return false
	}

	field := fl.Parent().FieldByName(
		strings.Replace(
			cases.Title(language.English).String(fieldName), " ", "", -1,
		),
	)
	if !field.IsValid() {
		return false
	}

	fieldParse, err := time.Parse(layout, field.String())
	if err != nil {
		log.Error(err)
		return false
	}

	if parse.Before(fieldParse) {
		return false
	}

	return true
}
