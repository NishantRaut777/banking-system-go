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
	accountIDStr := c.Param("id")

	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid account id",
		})
		return
	}

	// logged-in user_id from JWT
	userID := c.MustGet("user_id").(uuid.UUID)

	var req models.DepositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Deposit(
		c.Request.Context(),
		userID,
		accountID,
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
