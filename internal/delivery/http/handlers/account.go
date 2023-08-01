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

func (h AccountHandler) GetAllAccounts(c *gin.Context) {
	allAccounts, err := h.uc.GetAllAccounts(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, allAccounts)
}

func (h AccountHandler) GetMe(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

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
