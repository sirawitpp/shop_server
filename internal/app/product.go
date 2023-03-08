package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (p *productServer) GetProduct(ctx context.Context, req *pb.GetProductsReq) (*pb.Product, error) {
	result, err := p.productService.GetProduct(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pb.Product{
		ID:        result.ID,
		Name:      result.Name,
		Price:     float32(result.Price),
		Details:   result.Details,
		ImageUrl:  result.ImageUrl,
		CreatedAt: timestamppb.New(result.CreatedAt),
	}, nil
}

func (p *productServer) GetProducts(ctx context.Context, req *pb.GetProductsReq) (*pb.GetProductsRes, error) {
	result, err := p.productService.GetProducts(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	var products []*pb.Product
	for i := 0; i < len(result); i++ {
		products = append(products, &pb.Product{
			ID:        result[i].ID,
			Name:      result[i].Name,
			Price:     float32(result[i].Price),
			Details:   result[i].Details,
			ImageUrl:  result[i].ImageUrl,
			CreatedAt: timestamppb.New(result[i].CreatedAt),
		})
	}
	return &pb.GetProductsRes{
		Products: products,
	}, nil
}

func (p *productServer) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	username, err := p.getUsernameFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	if username != "sirawit23" {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	result, err := p.productService.CreateProduct(model.Product{
		Name:     req.GetName(),
		Details:  req.GetDetails(),
		Price:    float64(req.GetPrice()),
		ImageUrl: req.GetImageUrl(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.Product{
		ID:        result.ID,
		Name:      result.Name,
		Price:     float32(result.Price),
		Details:   result.Details,
		ImageUrl:  result.ImageUrl,
		CreatedAt: timestamppb.New(result.CreatedAt),
	}, nil

}
