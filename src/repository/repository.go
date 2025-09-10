
package repository

import (
	"context"
	"tracker/entity"
)

// TransactionRepository defines the interface for transaction repository operations
type TransactionRepository interface {
	Add(ctx context.Context, u entity.Transaction) error
	Edit(ctx context.Context, u entity.Transaction) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
	FindById(ctx context.Context, id int) (entity.Transaction, error)
}
