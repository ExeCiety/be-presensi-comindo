package responses

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"

	"github.com/gofrs/uuid"
)

type UserForLogin struct {
	Id           uuid.UUID      `json:"id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Name         string         `json:"name"`
	Roles        []*models.Role `json:"roles"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
}

type FindUsers struct {
	Id        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Nik       string         `json:"nik"`
	Name      string         `json:"name"`
	Roles     []*models.Role `json:"roles" gorm:"many2many:role_users;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

type FindUser struct {
	Id        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Nik       string         `json:"nik"`
	Name      string         `json:"name"`
	Roles     []*models.Role `json:"roles" gorm:"many2many:role_users;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

type CreateUser struct {
	Id        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Nik       string         `json:"nik"`
	Name      string         `json:"name"`
	Roles     []*models.Role `json:"roles" gorm:"many2many:role_users;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
	CreatedAt string         `json:"created_at"`
}

type UpdateUser struct {
	Id        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Nik       string         `json:"nik"`
	Name      string         `json:"name"`
	Roles     []*models.Role `json:"roles" gorm:"many2many:role_users;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
	CreatedAt string         `json:"created_at"`
}

type DeleteUsers struct {
	Id uuid.UUID `json:"id"`
}
