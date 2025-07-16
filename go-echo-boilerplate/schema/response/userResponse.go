package response

import (
	"boilerplate/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Role      string             `json:"role"`
	IsActive  bool               `json:"is_active"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type GetUsersResponse struct {
	Users []*UserResponse `json:"users"`
	Meta  utils.Meta      `json:"meta,omitempty"`
}
