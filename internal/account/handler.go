package account

import (
	"net/http"

	"github.com/NishantRaut777/banking-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service AccountService
}

func NewHandler(service AccountService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMyAccounts(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	accounts, err := h.service.GetMyAccounts(c.Request.Context(), userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetAccount(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	accountID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	account, err := h.service.GetAccountByID(
		c.Request.Context(),
		userID,
		accountID,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) Deposit(c *gin.Context) {
	var req models.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uuid.UUID)

	if err := h.service.Deposit(
		c.Request.Context(),
		userID,
		req.AccountID,
		req.Amount,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit successful",
	})
}

func (h *Handler) Withdraw(c *gin.Context) {
	var req models.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uuid.UUID)

	err := h.service.Withdraw(
		c.Request.Context(),
		userID,
		req.AccountID,
		req.Amount,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
}
