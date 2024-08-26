package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/application/usecases"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/data/repositories"
)

type MerchantService struct {
	Repo                  *repositories.MerchantRepository
	Validator             *validator.ValidatorAdapter
	CreateMerchantUseCase *usecases.CreateMerchantUseCase
}

func NewMerchantService(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *MerchantService {
	repo := repositories.NewMerchantRepository(db)
	return &MerchantService{
		Repo:                  repo,
		Validator:             validator,
		CreateMerchantUseCase: usecases.NewCreateMerchantUseCase(repo),
	}
}

// CreateMerchant godoc
// @Summary Create a merchant
// @Description Create a merchant
// @Tags merchant
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {string} string "Merchant created"
// @Router /merchant [post]
func (s *MerchantService) CreateMerchant(c *gin.Context) {
	// get user claim from JWT
	userID, _ := c.Get("userID")
	userIDStr := fmt.Sprintf("%v", userID)

	err := s.CreateMerchantUseCase.Execute(userIDStr)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Merchant created",
	})
}
