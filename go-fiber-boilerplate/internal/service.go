package internal

import "gitlab.com/sugaanaluam/gofiber-boilerplate/internal/service"

type ServiceManager struct {
	product service.ProductService
}

func NewServiceManager(repo RepositoryManager) ServiceManager {

	return ServiceManager{
		product: service.NewProductService(repo.product),
	}
}
