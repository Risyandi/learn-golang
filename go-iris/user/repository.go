package user

import "go-iris/pkg/database"

type Repository interface { // Repo methods here...
}

type repo struct { // Hold database instance here: e.g.
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repo{db: db}
}
