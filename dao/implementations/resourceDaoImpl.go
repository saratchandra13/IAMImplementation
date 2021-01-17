package implementations

import "samplego/models"

type ResourceDaoImpl struct {

}

func (resourceDaoImpl ResourceDaoImpl) GetResourceById(id int64) models.Resource {
	return models.AllResources[id]
}

