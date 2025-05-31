package repository

import (
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Add(model.User) error

	GetOne(email string) (model.User, error)

	UpdateUser(userId uuid.UUID, newData UpdateUserParams) error

	Delete(userId uuid.UUID) error
}

type UpdateUserParams struct {
	NewPassword string
}
