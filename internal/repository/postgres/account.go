package postgres

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"notification/internal/models"
)

type AccountRepository interface {
	GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)

	DeleteAccount(ctx context.Context, id uuid.UUID) error

	UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}
