package jwt

import (
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	ID        string          `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Email     string          `json:"email"`
	Role      model.Role      `json:"role"`
	ExpiresAt jwt.NumericDate `json:"expires_at"`
}
