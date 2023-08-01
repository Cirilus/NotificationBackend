package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"notification/internal/models"
	"notification/internal/usecase"
)

type AccountHandler struct {
	uc     *usecase.AccountUsecase
	logger *logrus.Logger
}

func NewHandler(uc *usecase.AccountUsecase, logger *logrus.Logger) *AccountHandler {
	return &AccountHandler{uc: uc, logger: logger}
}

// GetAccountById godoc
// @Summary Return account by id
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Success 201 {object} models.Account
// @Failure 500
// @Router /api/account/{id} [get]
func (h AccountHandler) GetAccountById(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("Cannot convert id as uuid = %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	accountById, err := h.uc.GetAccountById(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accountById)
}

// GetAllAccounts godoc
// @Summary Return all accounts
// @Tags account
// @Accept  json
// @Produce json
// @Success 201 {array} models.Account
// @Failure 500
// @Router /api/account/ [get]
func (h AccountHandler) GetAllAccounts(c *gin.Context) {
	allAccounts, err := h.uc.GetAllAccounts(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, allAccounts)
}

//// GetMe godoc
//func (h AccountHandler) GetMe(c *gin.Context) {
//	//TODO implement me
//	panic("implement me")
//}

// DeleteAccount godoc
// @Summary DeleteAccount
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Success 201
// @Failure 500
// @Router /api/account/{id} [delete]
func (h AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("Cannot convert id as uuid = %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteAccount(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// UpdateAccount godoc
// @Summary Update the Account
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Param account body models.Account true "Account object that needs to be updated"
// @Success 201
// @Failure 500
// @Router /api/account/{id} [put]
func (h AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("Cannot convert id as uuid = %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	inp := new(models.Account)
	err = c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request = %s", err)
		return
	}

	updateAccount, err := h.uc.UpdateAccount(c.Request.Context(), accountUUID, *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, updateAccount)
}

// CreateAccount godoc
// @Summary Create the Account
// @Tags account
// @Accept  json
// @Produce json
// @Param account body models.Account true "Account object that needs to be created"
// @Success 201
// @Failure 500
// @Router /api/account/ [post]
func (h AccountHandler) CreateAccount(c *gin.Context) {
	inp := new(models.Account)
	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request = %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.uc.CreateAccount(c.Request.Context(), *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
