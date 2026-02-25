package contracts

import (
	"context"

	"github.com/alazarbeyenenew2/shop/internal/app/product/domain"
)

type ProductRepository interface {
	Execute(ctx context.Context, p *domain.Product) (string, error)
	// FindByID(ctx context.Context, id string) (*domain.Product, error)
	// ListActive(ctx context.Context, categoryID string, limit, offset int) ([]*domain.Product, error)
}

type PorductService interface {
	CreateProduct(ctx context.Context, p *domain.Product) (string, error)
}
