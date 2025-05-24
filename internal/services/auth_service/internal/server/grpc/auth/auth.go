package auth

import (
	"context"

	pb "github.com/alexey-dobry/rental-service/internal/pkg/gen/auth"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) Register(ctx context.Context, req *pb.RegisterRequest) (pb.RegisterResponse, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}

func (s *ServerAPI) Login(ctx context.Context, req *pb.LoginRequest) (pb.LoginResponse, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}

func (s *ServerAPI) Auth(ctx context.Context, req *pb.AuthRequest) (emptypb.Empty, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}

func (s *ServerAPI) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (emptypb.Empty, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}

func (s *ServerAPI) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (pb.RefreshTokenResponse, error) {
	panic("AAAAAAAAAAAAAAAAAAA implement me")
}
