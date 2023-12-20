package auth

import (
	"errors"

	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type UserAuthed struct {
	Id       uuid.UUID          `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Roles    []*userModels.Role `json:"roles"`
}

var UserAuthedData *UserAuthed

func SetUserAuthed(c *fiber.Ctx) error {
	user := GetUserAuthedFromContext(c)
	claims := user.Claims.(jwt.MapClaims)

	roles := make([]*userModels.Role, 0)
	for _, role := range claims["roles"].([]interface{}) {
		roles = append(roles, &userModels.Role{
			Id:          uuid.FromStringOrNil(role.(map[string]interface{})["id"].(string)),
			Name:        role.(map[string]interface{})["name"].(string),
			RoleName:    role.(map[string]interface{})["role_name"].(string),
			Description: role.(map[string]interface{})["description"].(string),
		})
	}

	UserAuthedData = &UserAuthed{
		Id:       uuid.FromStringOrNil(claims["id"].(string)),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		Roles:    roles,
	}

	return nil
}

func IsUserAuthed() bool {
	if UserAuthedData == nil {
		return false
	}

	return true
}

func GetUserAuthedFromContext(c *fiber.Ctx) *jwt.Token {
	if c.Locals("user") == nil {
		return nil
	}

	return c.Locals("user").(*jwt.Token)
}

func GetUserData() *UserAuthed {
	return UserAuthedData
}

func IsUserAuthedHasRoles(roles []string) (bool, error) {
	if UserAuthedData == nil {
		return false, errors.New(utilsEnums.ErrorUserIsNotAuthenticated)
	}

	for _, role := range UserAuthedData.Roles {
		for _, roleName := range roles {
			if role.RoleName == roleName {
				return true, nil
			}
		}
	}

	return false, nil
}
