package repository

import entity "househonterbackend/v2/Entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
