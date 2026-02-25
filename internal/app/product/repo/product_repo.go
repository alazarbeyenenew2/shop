package repo

import (
	"context"
	"math/big"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/alazarbeyenenew2/shop/internal/app/product/domain"
	"go.uber.org/zap"
)

type ProductRepository struct {
	client *spanner.Client
	logger *zap.Logger
}

func NewProductRepository(
	client *spanner.Client,
	logger *zap.Logger,
) *ProductRepository {

	return &ProductRepository{
		client: client,
		logger: logger,
	}
}
func (r *ProductRepository) Execute(ctx context.Context, p *domain.Product) (string, error) {
	var mutation *spanner.Mutation
	r.logger.Info("saving product", zap.String("product_id", p.ID))

	_, err := r.client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {

		mutation = spanner.InsertOrUpdate(
			"products",
			[]string{
				"product_id",
				"name",
				"description",
				"category",
				"base_price_numerator",
				"base_price_denominator",
				"discount_percent",
				"discount_start_date",
				"discount_end_date",
				"status",
				"created_at",
				"updated_at",
			},
			[]interface{}{
				p.ID,
				p.Name,
				p.Description,
				p.CategoryID,
				p.BasePriceNumerator,
				p.BasePriceDenominator,
				toNullNumeric(p.DiscountPercent),
				p.DiscountStartDate,
				p.DiscountEndDate,
				p.Status,
				p.CreatedAt,
				p.UpdatedAt,
			},
		)

		if err := txn.BufferWrite([]*spanner.Mutation{mutation}); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		r.logger.Error("failed saving product",
			zap.Error(err),
			zap.String("product_id", p.ID),
		)
	}

	return p.ID, err
}

func toNullNumeric(value string) spanner.NullNumeric {
	value = strings.TrimSpace(value)

	if value == "" {
		return spanner.NullNumeric{Valid: false}
	}

	r := new(big.Rat)
	_, ok := r.SetString(value)
	if !ok {
		return spanner.NullNumeric{Valid: false}
	}

	return spanner.NullNumeric{
		Numeric: *r,
		Valid:   true,
	}
}
