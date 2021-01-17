package interfaces

import "samplego/models"

type ResourceDao interface {
	GetResourceById(id int64) models.Resource
}
