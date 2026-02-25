package product

import (
	"context"

	"github.com/alazarbeyenenew2/shop/internal/app/product/contracts"
	"github.com/alazarbeyenenew2/shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	ProductService contracts.PorductService
}

func (p *ProductHandler) UpdateProduct(context.Context, *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (p *ProductHandler) ActivateProduct(context.Context, *pb.ActivateProductRequest) (*pb.ActivateProductReply, error) {
	return nil, status.Error(codes.Unimplemented, "method ActivateProduct not implemented")
}
func (p *ProductHandler) DeactivateProduct(context.Context, *pb.DeactivateProductRequest) (*pb.DeactivateProductReply, error) {
	return nil, status.Error(codes.Unimplemented, "method DeactivateProduct not implemented")
}
func (p *ProductHandler) GetProduct(context.Context, *pb.GetProductRequest) (*pb.GetProductReply, error) {
	return nil, status.Error(codes.Unimplemented, "method GetProduct not implemented")
}
func (p *ProductHandler) ListProducts(context.Context, *pb.ListProductsRequest) (*pb.ListProductsReply, error) {
	return nil, status.Error(codes.Unimplemented, "method ListProducts not implemented")
}
