package handler

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/request"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/response"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/service"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/helper"
)

type ProductHandler interface {
	GetAllProducts(ctx *fiber.Ctx) error
	GetAllProductsWithParam(ctx *fiber.Ctx) error
	GetProductbyIDHandler(ctx *fiber.Ctx) error
	CreateProductsHandler(ctx *fiber.Ctx) error
	UpdateProductsHandler(ctx *fiber.Ctx) error
	DeleteProductsHandler(ctx *fiber.Ctx) error
}

type productHandler struct {
	service   service.ProductService
	conf      *config.Config
	exception exception.Exception
	validator *helper.XValidator
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &productHandler{
		service:   service,
		conf:      config.Get(),
		exception: exception.NewException("ProductHandler"),
		validator: helper.NewValidator(),
	}
}

func (h productHandler) GetAllProducts(ctx *fiber.Ctx) error {
	data := h.service.GetAllProductService(ctx.Context())

	return response.OK(ctx, constants.ResponseSuccessGet, data)
}

func (h productHandler) GetAllProductsWithParam(ctx *fiber.Ctx) error {
	req := new(request.GetProductRequest)
	req.Sort = ctx.Query("sort")
	req.Limit = ctx.Query("limit")

	err := h.validator.ValidateRequest(ctx, req)
	h.exception.BadRequest(err)

	data := h.service.GetAllProductService(ctx.Context())

	return response.OK(ctx, constants.ResponseSuccessGet, data)
}

func (h productHandler) GetProductbyIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data := h.service.GetProductbyIDService(ctx.Context(), id)

	return response.OK(ctx, constants.ResponseSuccessGet, data)
}

func (h productHandler) CreateProductsHandler(ctx *fiber.Ctx) error {
	req := new(request.CreateProductRequest)

	err := ctx.BodyParser(&req)
	h.exception.Error(err)

	err = h.validator.ValidateRequest(ctx, req)
	h.exception.BadRequest(err)

	data := h.service.CreateProductService(ctx.Context(), req)

	return response.Created(ctx, constants.ResponseSuccessCreate, data)
}

func (h productHandler) UpdateProductsHandler(ctx *fiber.Ctx) error {
	req := new(request.UpdateProductRequest)
	id := ctx.Params("id")
	err := ctx.BodyParser(&req)
	h.exception.Error(err)
	req.ID = id
	err = h.validator.ValidateRequest(ctx, req)
	h.exception.BadRequest(err)

	data := h.service.UpdateProductService(ctx.Context(), req)

	return response.OK(ctx, constants.ResponseSuccessUpdate, data)
}

func (h productHandler) DeleteProductsHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := h.service.DeleteProductService(ctx.Context(), id)
	h.exception.Error(err)

	return response.OK(ctx, constants.ResponseSuccessDelete, nil)
}
