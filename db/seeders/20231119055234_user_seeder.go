package seeders

import (
	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"
	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func UserSeeder() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231119055234",
		Migrate: func(tx *gorm.DB) error {
			var roles []*userModels.Role
			var adminRole *userModels.Role
			var hrdRole *userModels.Role
			var employeeRole *userModels.Role

			tx.Model(&userModels.Role{}).
				Where("role_name in ('superadmin', 'admin', 'hrd', 'employee')").
				Find(&roles)

			for i, v := range roles {
				switch v.RoleName {
				case userEnums.RoleNameAdmin:
					adminRole = roles[i]
				case userEnums.RoleNameHrd:
					hrdRole = roles[i]
				case userEnums.RoleNameEmployee:
					employeeRole = roles[i]
				}
			}

			admin01Password, _ := utils.HashPassword("admin01")
			hrd01Password, _ := utils.HashPassword("hrd01")
			employee01Password, _ := utils.HashPassword("employee01")

			return tx.Debug().
				Model(&userModels.User{}).Create(&[]userModels.User{
				{
					Username: "admin01",
					Email:    "admin01@comindo.com",
					Nik:      "100000000",
					Password: admin01Password,
					Name:     "Admin 01",
					Roles:    []*userModels.Role{adminRole},
				},
				{
					Username: "hrd01",
					Email:    "hrd01@comindo.com",
					Nik:      "100000001",
					Password: hrd01Password,
					Name:     "HRD 01",
					Roles:    []*userModels.Role{hrdRole},
				},
				{
					Username: "employee01",
					Email:    "employee01@comindo.com",
					Nik:      "100000002",
					Password: employee01Password,
					Name:     "Employee 01",
					Roles:    []*userModels.Role{employeeRole},
				},
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Unscoped().
				Delete(&userModels.User{}, "username in ('admin01', 'hrd01', 'employee01')").Error
		},
	}
}
