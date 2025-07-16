package constant

const (
	// HTTP Status Messages
	StatusSuccess = "success"
	StatusError   = "error"

	// Authentication
	AuthTokenKey     = "token"
	AuthUserKey      = "user"
	AuthHeaderKey    = "Authorization"
	AuthBearerPrefix = "Bearer "

	// Database Collections
	UsersCollection    = "users"
	ProductsCollection = "products"

	// Cache Keys
	UserCachePrefix    = "user:"
	ProductCachePrefix = "product:"

	// User Roles
	RoleAdmin = "admin"
	RoleUser  = "user"

	// JWT Claims
	JWTUserID = "user_id"
	JWTEmail  = "email"
	JWTRole   = "role"

	// Pagination
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100
)
