package services

import (
	"goravel/app/models"
	transactionRepository "goravel/app/repositories"
	"time"
)

type TransactionService struct {
	transactionRepository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (r *TransactionService) Create(clientId string, tipo string, description string, value string) (models.Transaction, error) {
	transaction := models.Transaction{
		ClientId:    clientId,
		Type:        tipo,
		Description: description,
		Value:       value,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	repository := transactionRepository.NewTransactionRepository()
	response, err := repository.Create(transaction)
	return response, err
}

func (r *TransactionService) GetLastTransactions(id int) []models.Transaction {
	repository := transactionRepository.NewTransactionRepository()
	response := repository.GetLastTenTransactionsByClientId(id)
	return response
}
