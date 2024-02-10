package repositories

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type ClientRepository struct {
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}

func (r *ClientRepository) GetById(id string) models.Client {
	var client models.Client
	facades.Orm().Query().Where("id", id).Get(&client)
	return client
}
