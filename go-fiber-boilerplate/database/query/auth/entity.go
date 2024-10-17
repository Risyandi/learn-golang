package auth

import "github.com/google/uuid"

type (
	UserStatusEntity struct {
		UserID      uuid.UUID `json:"userId"`
		IsActive    bool      `json:"isActive"`
		IsSuspended bool      `json:"isSuspended"`
		IsKyc       bool      `json:"isKyc"`
	}
)
