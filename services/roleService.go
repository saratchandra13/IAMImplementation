package services

import (
	"math/rand"
	dao "samplego/dao/implementations"
	"samplego/models"
)

type RoleService struct{

}

type RoleServiceInterface interface {
	AddRoleToUser(RoleId int64, UserId int64) models.User
	CreateRole(resourceId int64, Permissions []string, Name string) models.Role
}

func (roleService *RoleService) AddRoleToUser(RoleId int64, UserId int64) models.User{
	var userImpl dao.UserImpl

	user := userImpl.GetUserById(UserId)
	user.RoleIds = append(user.RoleIds, int(RoleId))
	userImpl.SaveUser(user)
	return user
}

func (roleService *RoleService) CreateRole(resourceId int64, Permissions []string, Name string) models.Role{
	var resourceImpl dao.ResourceDaoImpl
	var roleImpl dao.RoleDaoImpl
	resource := resourceImpl.GetResourceById(resourceId)
	if resource != (models.Resource{}) {
		role := models.Role{
			Id:          rand.Intn(10),
			Name:        Name,
			Permissions: map[int64][]string{resourceId: Permissions},
		}
		return roleImpl.AddRole(role)
	}else{
		return models.Role{}
	}
}

func (roleService *RoleService) CheckIfUserIsEligibleForAnOperation(resourceId int64, operation string, userId int64) bool {
	var userImpl dao.UserImpl
	var roleImpl dao.RoleDaoImpl
	user := userImpl.GetUserById(userId)
	for _, roleId := range user.RoleIds {
		role := roleImpl.GetRoleById(int64(roleId))
		for rsId, perms := range role.Permissions{
			if resourceId == rsId {
				if contains(perms,operation) {
					return true
				}
			}
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}