package implementations

import "samplego/models"

type UserImpl struct {

}

func (userImpl UserImpl) GetUserById(id int64) models.User {
	return models.AllUsers[id]
}

func  (userImpl UserImpl) SaveUser(user models.User) bool {
	models.AllUsers[user.Id] = user
	return true
}