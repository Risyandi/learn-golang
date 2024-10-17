package product

import (
	"gopkg.in/guregu/null.v4/zero"
)

type GetDataEntity struct {
	ID          string      `json:"Id" db:"id"`
	Name        zero.String `json:"Name" db:"name"`
	Description zero.String `json:"Description" db:"description"`
	CreatedAt   zero.Time   `json:"CreatedAt" db:"created_at"`
	UpdatedAt   zero.Time   `json:"UpdatedAt" db:"updated_at"`
	DeletedAt   zero.Time   `json:"DeletedAt" db:"deleted_at"`
	DeletedBy   zero.String `json:"DeletedBy" db:"deleted_by"`
}
