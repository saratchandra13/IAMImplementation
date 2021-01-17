package interfaces

import "samplego/models"

type UserDao interface {
	GetUserById(id int64) models.User
	SaveUser(user models.User) bool
}

