package request

type (
	GetProductRequest struct {
		Search       string `query:"search" validate:"omitempty,ascii"`                                                                      // search string
		ProductId    string `query:"productId" validate:"omitempty,uuid4"`                                                                   // product id
		UserId       string `query:"userId" validate:"omitempty,uuid4"`                                                                      // user id
		Status       string `query:"status" validate:"omitempty,oneof=100 200 300 400 500 600"`                                              // status of product
		Type         string `query:"type" validate:"omitempty,oneof=100 200"`                                                                // type of product
		Category     string `query:"category" validate:"omitempty,number"`                                                                   // category of product
		Favorite     string `query:"favorite" validate:"omitempty,oneof=true false"`                                                         // favorite of product
		ExtProductId string `query:"extProductId" validate:"omitempty,ascii"`                                                                // external product id of product
		ExtShopId    string `query:"extShopId" validate:"omitempty,ascii"`                                                                   // external shop id of product
		ExtSourceId  string `query:"extSourceId" validate:"omitempty,oneof=100 200 300 400 500"`                                             // external source id of product
		Sort         string `query:"sort" validate:"omitempty,oneof=created_at name price_unit price_discounted price_wholesale price_cogs"` // sort of product
		Order        string `query:"order" validate:"omitempty,oneof=asc desc"`                                                              // order string
		Limit        string `query:"limit" validate:"omitempty,number"`                                                                      // limit string
		Offset       string `query:"offset" validate:"omitempty,number"`                                                                     // offset string
	}

	CreateProductRequest struct {
		ID          string `json:"-"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"omitempty"`
	}

	UpdateProductRequest struct {
		ID          string `json:"Id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"omitempty"`
	}
)
