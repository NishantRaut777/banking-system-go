package account

import (
	"net/http"

	"github.com/NishantRaut777/banking-system-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service AccountService
}

func NewHandler(service AccountService) *Handler {
	return &Handler{service: service}
}

// GetMyAccount godoc
// @Summary Get logged-in user's account
// @Description Returns account details of logged-in user
// @Tags Transaction
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.GetMyAccount
// @Failure 401 {object} dto.GetMyAccountUnathorised
// @Failure 500 {object} dto.GetMyAccountServerError
// @Router /accounts [get]
func (h *Handler) GetMyAccount(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	accounts, err, statusCode := h.service.GetMyAccount(c.Request.Context(), userID)

	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
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

// Deposit godoc
// @Summary Deposit money
// @Description Deposit money into logged-in user's account
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.DepositRequest true "Deposit payload"
// @Success 200 {object} dto.DepositResponse
// @Failure 401 {object} dto.DepositResponseClientError
// @Failure 500 {object} dto.TransferResponseServerError
// @Router /accounts/deposit [post]
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

// Withdraw godoc
// @Summary Withdraw money
// @Description Withdraw money from logged-in user's account
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.WithdrawRequest true "Withdraw payload"
// @Success 200 {object} dto.WithdrawResponse
// @Failure 401 {object} dto.WithdrawResponseClientError
// @Failure 500 {object} dto.WithdrawResponseServerError
// @Router /accounts/withdraw [post]
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

// Transfer godoc
// @Summary Transfer money
// @Description Transfer money to another account
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.TransferRequest true "Transfer payload"
// @Success 200 {object} dto.TransferResponse
// @Failure 401 {object} dto.TransferResponseClientError
// @Failure 500 {object} dto.TransferResponseServerError
// @Router /accounts/transfer [post]
func (h *Handler) Transfer(c *gin.Context) {
	var req models.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.FromAccountID == req.ToAccountID {
		c.JSON(400, gin.H{"error": "same account transfer not allowed"})
		return
	}

	userID := c.MustGet("user_id").(uuid.UUID)

	err := h.service.Transfer(
		c.Request.Context(),
		userID,
		req.FromAccountID,
		req.ToAccountID,
		req.Amount,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "transfer successful"})
}
