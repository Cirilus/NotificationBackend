package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"notification/internal/models"
	"notification/internal/usecase"
)

type AccountHandler struct {
	uc *usecase.AccountUsecase
}

func NewHandler(uc *usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{uc: uc}
}

func (h AccountHandler) GetAccountById(c *gin.Context, id uuid.UUID) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (h AccountHandler) GetAllAccounts(c *gin.Context) ([]models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (h AccountHandler) GetMe(c *gin.Context) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (h AccountHandler) DeleteAccount(c *gin.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (h AccountHandler) UpdateAccount(c *gin.Context, id uuid.UUID, account models.Account) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (h AccountHandler) CreateAccount(c *gin.Context, account models.Account) error {
	//TODO implement me
	panic("implement me")
}
