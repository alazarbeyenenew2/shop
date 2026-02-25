package initiator

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/alazarbeyenenew2/shop/internal/app/product/contracts"
	"github.com/alazarbeyenenew2/shop/internal/app/product/repo"
	"go.uber.org/zap"
)

type Repo struct {
	ProductRepo contracts.ProductRepository
}

func InitRepo(ctx context.Context, client *spanner.Client, logger *zap.Logger) Repo {
	return Repo{
		ProductRepo: repo.NewProductRepository(client, logger),
	}
}
