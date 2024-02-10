package services

import (
	"goravel/app/models"
	clientRepository "goravel/app/repositories"
)

type ClientService struct {
	clientRepository.ClientRepository
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (r *ClientService) GetById(id string) models.Client {
	repository := clientRepository.NewClientRepository()
	return repository.GetById(id)
}
