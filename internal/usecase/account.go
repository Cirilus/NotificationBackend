package usecase

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"notification/internal/models"
	"notification/internal/repository/postgres"
)

type AccountUsecaseInterface interface {
	GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)
	GetMe(ctx context.Context) (*models.Account, error)

	DeleteAccount(ctx context.Context, id uuid.UUID) error

	UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}

type AccountUsecase struct {
	repo postgres.AccountRepository
}

func NewAccountUsecase(repo postgres.AccountRepository) *AccountUsecase {
	return &AccountUsecase{repo: repo}
}

func (a AccountUsecase) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountUsecase) GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountUsecase) GetMe(ctx context.Context) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountUsecase) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountUsecase) UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountUsecase) CreateAccount(ctx context.Context, account models.Account) error {
	//TODO implement me
	panic("implement me")
}
