package internal

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
)

type RouterManager struct {
	v1         fiber.Router
	handler    HandlerManager
	middleware MiddlewareManager
	exception  exception.Exception // Example exception module
}

func NewRouterManager(route *fiber.App, handler HandlerManager, middleware MiddlewareManager) *RouterManager {
	conf := config.Get()
	api := route.Group(conf.API.Route)
	return &RouterManager{
		v1:         api.Group("v1"),
		handler:    handler,
		middleware: middleware,
		exception:  exception.NewException("route"), // Example exception initialization
	}
}

func (rm *RouterManager) Init() {
	rm.v1.Get("/exception", func(c *fiber.Ctx) error {
		rm.exception.BadRequest(errors.New("this is panic error"))
		return c.JSON("exception")
	})
	product := rm.v1.Group("/product", rm.middleware.auth.Auth(), rm.middleware.auth.ValidateToken(), rm.middleware.auth.ValidateAccount(), rm.middleware.session.CreateSession(), rm.middleware.session.LogSession())
	product.Post("/", rm.handler.product.CreateProductsHandler)
	product.Put("/:id", rm.handler.product.UpdateProductsHandler)
	product.Delete("/:id", rm.handler.product.DeleteProductsHandler)
	product.Get("/", rm.handler.product.GetAllProducts)
	product.Get("/param", rm.handler.product.GetAllProductsWithParam)
	product.Get("/:id", rm.handler.product.GetProductbyIDHandler)

}
