
package repository

import (
	"context"
	"tracker/entity"
	"github.com/jmoiron/sqlx"
)

// PGTransactionRepository defines the interface for transaction operations
type PGTransactionRepository interface {
	Add(ctx context.Context, u entity.Transaction) error
	Edit(ctx context.Context, u entity.Transaction) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
}

// pgTransactionRepository implements PGTransactionRepository
type pgTransactionRepository struct {
	DB *sqlx.DB
}

// NewPGTransactionRepository creates a new repository instance
func NewPGTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &pgTransactionRepository{
		DB: db,
	}
}

// FindAll retrieves all transactions from the database
func (r *pgTransactionRepository) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}
	query := `SELECT * FROM "transaction"`
	if err := r.DB.Select(&transactions, query); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *pgTransactionRepository) FindById(ctx context.Context, id int) (entity.Transaction, error) {
	transactions := &entity.Transaction{}
	query := `SELECT * FROM "transaction" WHERE id=$1`
	if err := r.DB.GetContext(ctx, transactions, query, id); err != nil {
		return *transactions, err
	}
	return *transactions, nil
}

// Add inserts a new transaction into the database
func (r *pgTransactionRepository) Add(ctx context.Context, u entity.Transaction) error {
	query := `INSERT INTO "transaction" (type, ticker, volume, price, date) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, u.Type, u.Ticker, u.Volume, u.Price, u.Date).Scan(&u.ID)
	return err
}

// Edit updates an existing transaction in the database
func (r *pgTransactionRepository) Edit(ctx context.Context, u entity.Transaction) error {
	query := `UPDATE "transaction" SET type=$1, ticker=$2, volume=$3, price=$4, date=$5 WHERE id=$6`
	_, err := r.DB.ExecContext(ctx, query, u.Type, u.Ticker, u.Volume, u.Price, u.Date, u.ID)
	return err
}

// Delete removes a transaction from the database by ID
func (r *pgTransactionRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM "transaction" WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
