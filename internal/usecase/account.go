package usecase

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	repo   *postgres.AccountRepository
	logger *logrus.Logger
}

func NewAccountUsecase(repo *postgres.AccountRepository, logger *logrus.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, logger: logger}
}

func (a AccountUsecase) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	accountById, err := a.repo.GetAccountById(ctx, id)
	if err != nil {
		a.logger.Error("Account - UseCase - AccountById")
		return nil, err
	}
	return accountById, nil
}

func (a AccountUsecase) GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	allAccounts, err := a.repo.GetAllAccounts(ctx)
	if err != nil {
		logrus.Error("Account - UseCase - AllAccounts")
		return nil, err
	}
	return allAccounts, nil
}

func (a AccountUsecase) GetMe(ctx context.Context) (*models.Account, error) {
	// TODO implement me
	panic("Implement me")
}

func (a AccountUsecase) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	err := a.repo.DeleteAccount(ctx, id)
	if err != nil {
		logrus.Error("Account - UseCase - DeleteAccount")
		return err
	}
	return nil
}

func (a AccountUsecase) UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error) {
	updateAccount, err := a.repo.UpdateAccount(ctx, id, account)
	if err != nil {
		logrus.Error("Account - UseCase - Account")
		return nil, err
	}
	return updateAccount, err
}

func (a AccountUsecase) CreateAccount(ctx context.Context, account models.Account) error {
	err := a.repo.CreateAccount(ctx, account)
	if err != nil {
		logrus.Error("Account - UseCase - CreateAccount")
		return err
	}
	return nil
}
