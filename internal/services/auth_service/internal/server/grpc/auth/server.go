package auth

import (
	pb "github.com/alexey-dobry/rental-service/internal/pkg/gen/auth"
	"github.com/alexey-dobry/rental-service/internal/pkg/logger"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/jwt"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/repository"
)

type ServerAPI struct {
	pb.UnimplementedAuthServer

	logger     logger.Logger
	repository repository.UserRepository
	jwtHandler jwt.JWTHandler
}

func NewGRPCServer(logger logger.Logger, repository repository.UserRepository) *ServerAPI {
	return &ServerAPI{
		repository: repository,
		logger:     logger.WithFields("layer", "grpc server api", "server", "manager"),
	}
}
