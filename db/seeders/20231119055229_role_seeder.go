package seeders

import (
	"fmt"

	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"
	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RoleSeeder() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231119055229",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				Model(&userModels.Role{}).Create(&[]userModels.Role{
				{
					Name:        "admin",
					RoleName:    userEnums.RoleNameAdmin,
					Description: "Role for admin",
				},
				{
					Name:        "HRD",
					RoleName:    userEnums.RoleNameHrd,
					Description: "Role for HRD",
				},
				{
					Name:        "employee",
					RoleName:    userEnums.RoleNameEmployee,
					Description: "Role for employee",
				},
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Unscoped().
				Delete(&userModels.Role{},
					fmt.Sprintf(
						`role_name in ('%s', '%s', '%s')`,
						userEnums.RoleNameAdmin, userEnums.RoleNameHrd, userEnums.RoleNameEmployee,
					),
				).Error
		},
	}
}
