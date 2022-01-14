package mongo

import (
	"context"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

//struct for communication with database
type FinanceRepo struct {
	db *mongo.Client
}

//new struct
func NewFinanceRepo(db *mongo.Client) *FinanceRepo {
	return &FinanceRepo{db: db}
}

//transaction from user
func (r *FinanceRepo) MakeTransaction(ctx context.Context, inp *domain.TransactionInput) error {
	return nil

}

//transaction from user to user
func (r *FinanceRepo) MakeRemittance(ctx context.Context, inp *domain.RemittanceInput) error {
	return nil
}

//user balance
func (r *FinanceRepo) GetBalance(ctx context.Context, inp *domain.BalanceInput) (float64, error) {
	return 0, nil
}

//list of all transactions  by query
func (r *FinanceRepo) GetTransactionsList(ctx context.Context, inp *domain.TransactionsListInput) ([]domain.TransactionsList, error) {
	return nil, nil
}

//create transaction
func (r *FinanceRepo) CreateNewTransaction(ctx context.Context, idFrom uuid.UUID, operation string, sum float64, idTo uuid.UUID, description string) error {
	return nil
}

//start migration
func (r *FinanceRepo) StartMigration(ctx context.Context, dir, dest string) error {
	return nil
}

//close db
func (r *FinanceRepo) Close(ctx context.Context) error {
	return r.db.Disconnect(ctx)
}
