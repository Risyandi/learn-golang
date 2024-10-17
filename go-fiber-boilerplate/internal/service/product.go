package service

import (
	"context"

	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/database/query/product"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/repository"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/request"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type ProductService interface {
	GetAllProductService(ctx context.Context) []product.GetDataEntity
	GetProductbyIDService(ctx context.Context, id string) product.GetDataEntity
	CreateProductService(ctx context.Context, req *request.CreateProductRequest) product.GetDataEntity
	UpdateProductService(ctx context.Context, req *request.UpdateProductRequest) product.GetDataEntity
	DeleteProductService(ctx context.Context, id string) error
}

type productService struct {
	repo      repository.ProductRepository
	conf      *config.Config
	log       zerolog.Logger
	exception exception.Exception
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo:      repo,
		conf:      config.Get(),
		log:       logger.Get("product-service"),
		exception: exception.NewException("product-service"),
	}
}

func (s productService) GetAllProductService(ctx context.Context) []product.GetDataEntity {
	data, err := s.repo.GetAllProductRepository(ctx)
	s.exception.Error(err)

	return data
}

func (s productService) GetProductbyIDService(ctx context.Context, id string) product.GetDataEntity {
	data, err := s.repo.GetProductbyIDRepository(ctx, id)
	s.exception.Error(err)

	return data
}

func (s productService) CreateProductService(ctx context.Context, req *request.CreateProductRequest) product.GetDataEntity {
	// Place your custom logic here
	tr, err := s.repo.NewTransaction(ctx)
	s.exception.Error(err)

	data, err := tr.CreateProductRepository(ctx, req)
	s.exception.Error(err)

	err = tr.Commit()
	s.exception.Error(err)

	return data
}

func (s productService) UpdateProductService(ctx context.Context, req *request.UpdateProductRequest) product.GetDataEntity {
	// Place your custom logic here

	tr, err := s.repo.NewTransaction(ctx)
	s.exception.Error(err)

	data, err := tr.UpdateProductRepository(ctx, req)
	s.exception.Error(err)

	err = tr.Commit()
	s.exception.Error(err)

	return data
}

func (s productService) DeleteProductService(ctx context.Context, id string) error {
	// Place your custom logic here
	tr, err := s.repo.NewTransaction(ctx)
	s.exception.Error(err)

	err = tr.DeleteProductRepository(ctx, id)
	s.exception.Error(err)

	err = tr.Commit()
	s.exception.Error(err)

	return nil
}
