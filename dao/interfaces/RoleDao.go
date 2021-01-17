package interfaces

import "samplego/models"

type RoleDao interface {
	GetRoleById(id int64) models.Role
	AddRole(role models.Role) models.Role
}

