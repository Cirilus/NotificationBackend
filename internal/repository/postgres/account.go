package postgres

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"notification/internal/models"
	"notification/pkg/clients/postgresql"
)

type AccountRepositoryInterface interface {
	GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)

	DeleteAccount(ctx context.Context, id uuid.UUID) error

	UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}

type AccountRepository struct {
	*postgresql.Postgres
	logger *logrus.Logger
}

func NewAccountRepository(postgres *postgresql.Postgres, logger *logrus.Logger) *AccountRepository {
	return &AccountRepository{Postgres: postgres, logger: logger}
}

func (a AccountRepository) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	sql := `SELECT id, telegram, email FROM account WHERE id = $1`
	account := new(models.Account)
	err := a.Pool.QueryRow(ctx, sql, id).Scan(&account.Id, &account.Telegram, &account.Email)
	if err != nil {
		a.logger.Error("Account - Repository - GetAccountById")
		return nil, err
	}
	return account, nil
}

func (a AccountRepository) GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	sql := `SELECT id, telegram, email FROM account`
	accounts := make([]models.Account, 0)
	rows, err := a.Pool.Query(ctx, sql)
	if err != nil {
		a.logger.Error("Account - Repository - GetAllAccounts")
		return nil, err
	}
	for rows.Next() {
		account := new(models.Account)
		err = rows.Scan(&account.Id, &account.Telegram, &account.Email)
		if err != nil {
			a.logger.Error("Account - Repository - GetAllAccounts")
			return nil, err
		}
		accounts = append(accounts, *account)
	}
	return accounts, nil
}

func (a AccountRepository) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM account WHERE id=$1`
	_, err := a.Pool.Exec(ctx, sql, id)
	if err != nil {
		a.logger.Error("Account - Repository - DeleteAccount")
		return err
	}
	return nil
}

func (a AccountRepository) UpdateAccount(ctx context.Context, id uuid.UUID, account models.Account) (*models.Account, error) {
	sql := `UPDATE account 
	SET id=COALESCE($2, id), telegram=COALESCE($3, telegram), 
	email=COALESCE($4, email)
	WHERE id = $1
	RETURNING id, telegram, email`

	updateAccount := new(models.Account)

	err := a.Pool.QueryRow(ctx, sql, id, account.Id, account.Telegram, account.Email).
		Scan(&updateAccount.Id, &updateAccount.Telegram, &updateAccount.Email)
	if err != nil {
		a.logger.Error("Account - Repository - Account")
		return nil, err
	}
	return updateAccount, nil
}

func (a AccountRepository) CreateAccount(ctx context.Context, account models.Account) error {
	sql := "INSERT INTO account(id, telegram, email) VALUES ($1, $2, $3)"
	_, err := a.Pool.Exec(ctx, sql, account.Id, account.Telegram, account.Email)
	if err != nil {
		a.logger.Error("Account - Repository - CreateAccount")
		return err
	}
	return nil
}
