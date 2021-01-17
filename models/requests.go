package models

type CreateRoleRequest struct {
	ResourceId int64 `json:"resource_id"`
	Perms []string `json:"perms"`
	RoleName string `json:"role_name"` 
}
