package model

import (
	"github.com/cyber-lama/personal-notes/api/internal/model/user"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(*sqlx.DB) (*user.User, error)
	Find(*sqlx.DB, int) (*user.User, error)
	FindByEmail(*sqlx.DB, string) (*user.User, error)
}
