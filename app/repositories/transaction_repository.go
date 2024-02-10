package repositories

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	error := facades.Orm().Query().Save(&transaction)
	return transaction, error
}

func (r *TransactionRepository) GetLastTenTransactionsByClientId(id int) []models.Transaction {
	var transactions []models.Transaction
	facades.Orm().Query().Where("client_id", id).Limit(10).Get(&transactions)
	return transactions
}
