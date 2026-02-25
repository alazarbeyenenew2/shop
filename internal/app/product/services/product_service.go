package services

import (
	"context"
	"time"

	"github.com/alazarbeyenenew2/shop/internal/app/product/contracts"
	"github.com/alazarbeyenenew2/shop/internal/app/product/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ProductService struct {
	repo   contracts.ProductRepository
	logger *zap.Logger
}

func InitProductService(repo contracts.ProductRepository,
	logger *zap.Logger) contracts.PorductService {
	return &ProductService{
		repo:   repo,
		logger: logger,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, p *domain.Product) (string, error) {

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Status = "ACTIVE"
	p.Version = 1
	p.ID = uuid.NewString()

	return s.repo.Execute(ctx, p)
}
