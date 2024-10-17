package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/database/query/product"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/request"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type ProductRepository interface {
	NewTransaction(ctx context.Context) (ProductRepositoryTx, error)
	GetAllProductRepository(ctx context.Context) ([]product.GetDataEntity, error)
	GetProductbyIDRepository(ctx context.Context, id string) (product.GetDataEntity, error)
}

type productRepository struct {
	db        *sql.DB
	exception exception.Exception
	log       zerolog.Logger
}

func NewProductRepository(pg *sql.DB) ProductRepository {
	return &productRepository{
		db:        pg,
		exception: exception.NewException("product-repository"),
		log:       logger.Get("product-repository"),
	}
}

type ProductRepositoryTx interface {
	Commit() error

	CreateProductRepository(ctx context.Context, req *request.CreateProductRequest) (product.GetDataEntity, error)
	UpdateProductRepository(ctx context.Context, req *request.UpdateProductRequest) (product.GetDataEntity, error)
	DeleteProductRepository(ctx context.Context, id string) error
}

type productRepositoryTx struct {
	ctx context.Context
	tx  *sql.Tx
}

func (r *productRepository) NewTransaction(ctx context.Context) (ProductRepositoryTx, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &productRepositoryTx{ctx, tx}, nil
}

func (r *productRepositoryTx) Commit() error {
	return r.tx.Commit()
}

func (r *productRepository) GetAllProductRepository(ctx context.Context) ([]product.GetDataEntity, error) {
	queryStr := product.GetProductSQL

	rows, err := r.db.QueryContext(ctx, queryStr)
	if err != nil {
		if err == sql.ErrNoRows {
			r.exception.ErrorWithoutNoSqlResult(err)
			return nil, err
		}
		return nil, err
	}
	defer rows.Close()

	var res []product.GetDataEntity
	for rows.Next() {

		var j product.GetDataEntity
		err := rows.Scan(
			&j.ID,
			&j.Name,
			&j.Description,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, j)
	}

	return res, nil
}

func (r productRepository) GetProductbyIDRepository(ctx context.Context, id string) (product.GetDataEntity, error) {
	queryStr := product.GetProductByIDSQL

	var res product.GetDataEntity

	err := r.db.QueryRowContext(ctx, queryStr, id).Scan(
		&res.ID,
		&res.Name,
		&res.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			r.exception.ErrorWithoutNoSqlResult(err)
			return res, err
		}
		return res, err
	}

	return res, nil
}

func (r productRepositoryTx) CreateProductRepository(ctx context.Context, req *request.CreateProductRequest) (product.GetDataEntity, error) {
	queryStr := product.CreateProductSQL

	var res product.GetDataEntity
	req.ID = uuid.New().String()
	err := r.tx.QueryRowContext(
		ctx,
		queryStr,
		req.ID,
		req.Name,
		req.Description,
		time.Now(),
	).Scan(
		&res.ID,
		&res.Name,
		&res.Description,
		&res.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		}
		return res, err
	}

	return res, nil
}

func (r productRepositoryTx) UpdateProductRepository(ctx context.Context, req *request.UpdateProductRequest) (product.GetDataEntity, error) {
	queryStr := product.UpdateProductSQL

	var res product.GetDataEntity

	err := r.tx.QueryRowContext(
		ctx,
		queryStr,
		req.Name,
		req.Description,
		time.Now(),
		req.ID,
	).Scan(
		&res.ID,
		&res.Name,
		&res.Description,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		}
		return res, err
	}

	return res, nil
}

func (r productRepositoryTx) DeleteProductRepository(ctx context.Context, id string) error {
	queryStr := product.DeleteProductSQL

	_, err := r.tx.ExecContext(ctx, queryStr, time.Now(), id, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}

	return nil
}
