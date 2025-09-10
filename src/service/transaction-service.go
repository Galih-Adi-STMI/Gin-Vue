package service

import (
	"context"
	"strconv"
	"tracker/entity"
	"tracker/repository"
)

// TransactionService defines the interface for transaction service operations
type TransactionService interface {
	Add(ctx context.Context, u entity.Transaction) error
	Edit(ctx context.Context, u entity.Transaction) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
	FindById(ctx context.Context, id string) (entity.Transaction, error)
}

// transactionService implements TransactionService
type transactionService struct {
	repository repository.TransactionRepository
}

// NewTransactionService creates a new service instance
func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{
		repository: repository,
	}
}

// FindAll retrieves all transactions via the repository
func (s *transactionService) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	return s.repository.FindAll(ctx)
}

func (s *transactionService) FindById(ctx context.Context, id string) (entity.Transaction, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return entity.Transaction{}, err
	}
	transaction, err := s.repository.FindById(ctx, i)
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}	


// Add creates a new transaction via the repository
func (s *transactionService) Add(ctx context.Context, transaction entity.Transaction) error {
	return s.repository.Add(ctx, transaction)
}

// Edit updates an existing transaction via the repository
func (s *transactionService) Edit(ctx context.Context, transaction entity.Transaction) error {
	return s.repository.Edit(ctx, transaction)
}

// Delete removes a transaction by ID via the repository
func (s *transactionService) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
