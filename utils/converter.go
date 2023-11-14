package utils

import (
	"strconv"

	"github.com/gofrs/uuid"
)

// StrToInt converts string number to int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// IntToStr converts int to string
func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// StrToPtr converts string to string pointer
func StrToPtr(str string) *string {
	return &str
}

// BoolToPtr converts bool to bool pointer
func BoolToPtr(b bool) *bool {
	return &b
}

// UUIDToPtr converts uuid to uuid pointer
func UUIDToPtr(id uuid.UUID) *uuid.UUID {
	return &id
}

// Float64ToPtr converts float64 to float64 pointer
func Float64ToPtr(f float64) *float64 {
	return &f
}
