package product

import (
	"context"
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/alazarbeyenenew2/shop/internal/app/product/domain"
	"github.com/alazarbeyenenew2/shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	if err := validateCreateRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	appReq := mapToCreateProductRequest(req)
	productID, err := h.ProductService.CreateProduct(ctx, appReq)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductReply{
		ProductId: productID,
	}, nil

}

func mapToCreateProductRequest(req *pb.CreateProductRequest) *domain.Product {
	var denominator int64
	var numerator int64
	if req == nil {
		return nil
	}

	str := strconv.FormatFloat(req.BasePrice, 'f', -1, 64)

	if !strings.Contains(str, ".") {
		numerator = int64(req.BasePrice)
		denominator = 1
	} else {
		parts := strings.Split(str, ".")
		decimalPlaces := len(parts[1])

		denominator = int64(math.Pow10(decimalPlaces))
		numerator = int64(math.Round(req.BasePrice * float64(denominator)))
	}

	return &domain.Product{
		Name:                 req.Name,
		Description:          req.Description,
		CategoryID:           req.Category,
		BasePriceNumerator:   numerator,
		BasePriceDenominator: denominator,
	}
}

func validateCreateRequest(req *pb.CreateProductRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if req.BasePrice <= 0 {
		return errors.New("price must be greater than zero")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	return nil
}
