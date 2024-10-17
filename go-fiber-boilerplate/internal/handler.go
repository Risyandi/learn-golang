package internal

import "gitlab.com/sugaanaluam/gofiber-boilerplate/internal/handler"

type HandlerManager struct {
	product handler.ProductHandler
}

func NewHandlerManager(service ServiceManager) HandlerManager {

	return HandlerManager{
		product: handler.NewProductHandler(service.product),
	}
}
