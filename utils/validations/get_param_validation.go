package validations

import (
	"strings"
)

func GetDateGreaterThanFieldParam(validationError *ValidationError) string {
	paramValues := strings.SplitN(validationError.Param, ";", 2)
	if len(paramValues) <= 0 {
		return ""
	}

	return strings.Replace(strings.ToLower(paramValues[0]), " ", "_", -1)
}
