package auth

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/alexey-dobry/rental-service/internal/pkg/gen/auth"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/jwt"
	"github.com/alexey-dobry/rental-service/internal/services/auth_service/internal/domain/model"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.User{
		ID:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role.String(),
	}
	err := s.repository.Add(user)
	if err != nil {
		errMsg := fmt.Sprintf("Error adding new user to data: %s", err)
		s.logger.Errorf(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	accessToken, refreshToken, err := s.jwtHandler.GenerateJWTPair(jwt.Claims{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      model.Role(user.Role),
	})

	if err != nil {
		errMsg := fmt.Sprintf("Failed to generate token pair: %s", err)
		s.logger.Errorf(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	// TODO: add here profile service handler call as soon as it implemented

	response := pb.RegisterResponse{
		JwtAccessToken:  accessToken,
		JwtRefreshToken: refreshToken,
	}

	return &response, nil
}

func (s *ServerAPI) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.repository.GetOne(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "User entry with given credentials not found")
	} else if err != nil {
		errMsg := fmt.Sprintf("Failed to get user data from database: %s", err)
		s.logger.Errorf(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	accessToken, refreshToken, err := s.jwtHandler.GenerateJWTPair(jwt.Claims{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      model.Role(user.Role),
	})

	if err != nil {
		errMsg := fmt.Sprintf("Failed to generate token pair: %s", err)
		s.logger.Errorf(errMsg)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &pb.LoginResponse{
		JwtAccessToken:  accessToken,
		JwtRefreshToken: refreshToken,
	}, nil
}

func (s *ServerAPI) Auth(ctx context.Context, req *pb.AuthRequest) (*emptypb.Empty, error) {
	_, err := s.jwtHandler.ValidateJWT(req.JwtAccessToken, jwt.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Access denied")
	}

	return &emptypb.Empty{}, nil
}

func (s *ServerAPI) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (pb.RefreshTokenResponse, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}
