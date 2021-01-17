package implementations

import "samplego/models"

type RoleDaoImpl struct {

}

func (RoleDaoImpl RoleDaoImpl) GetRoleById(id int64) models.Role {
	return models.AllRoles[id]
}

func (RoleDaoImpl RoleDaoImpl) AddRole(role models.Role) models.Role {
	models.AllRoles[int64(role.Id)] = role
	return role
}