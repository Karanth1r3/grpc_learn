package auth

import (
	"context"

	ssov1 "github.com/Karanth1r3/proto3_learn/proto/sso"
	"google.golang.org/grpc"
)

// Struct to link generated code with go logic
type serverAPI struct {
	// To be able to run app without implementing all generated interfaces => can throw unimplemented message instead
	ssov1.UnimplementedAuthServer
}

// Registers handler
func Register(grpcs *grpc.Server) {
	ssov1.RegisterAuthServer(grpcs, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement login")
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement register")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement isadmin")
}
