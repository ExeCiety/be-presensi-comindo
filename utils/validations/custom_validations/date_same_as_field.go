package custom_validations

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func DateSameAsField(fl validator.FieldLevel) bool {
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

	layoutParse, errParseLayout := time.Parse(layout, fl.Field().String())
	if errParseLayout != nil {
		log.Error(errParseLayout)
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

	fieldParse, errFieldParse := time.Parse(layout, field.String())
	if errFieldParse != nil {
		log.Error(errFieldParse)
		return false
	}

	if layoutParse.Format("2006-01-02") != fieldParse.Format("2006-01-02") {
		return false
	}

	return true
}
