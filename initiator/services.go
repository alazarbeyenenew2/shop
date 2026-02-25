package initiator

import (
	"github.com/alazarbeyenenew2/shop/internal/app/product/contracts"
	"github.com/alazarbeyenenew2/shop/internal/app/product/services"
	"go.uber.org/zap"
)

type Services struct {
	ProductService contracts.PorductService
}

func InitServices(repo Repo, logger *zap.Logger) Services {
	return Services{
		ProductService: services.InitProductService(repo.ProductRepo, logger),
	}
}
