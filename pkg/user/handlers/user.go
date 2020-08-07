package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/database"
	"go.uber.org/zap"
)

type User struct {
	l  *zap.SugaredLogger
	db *database.Users
}

func NewUserHandler(logger *zap.SugaredLogger, db *database.Users) *User {
	return &User{
		l:  logger,
		db: db,
	}
}
