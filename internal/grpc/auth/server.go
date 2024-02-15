package auth

import (
	"context"

	ssov1 "github.com/Karanth1r3/proto3_learn/proto/sso"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)

	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

// Struct to link generated code with go logic
type serverAPI struct {
	// To be able to run app without implementing all generated interfaces => can throw unimplemented message instead
	ssov1.UnimplementedAuthServer
	auth Auth
}

// Registers handler with authentification interface
func Register(grpcs *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(grpcs, &serverAPI{auth: auth})
}

// If id of user is 0 => throw error; for readability
const (
	emptyValue = 0
)

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {

	// Validation of request fields
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	// Referring to service layer (interface) to login operation
	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		// TODO - add error differentiation
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.LoginResponse{

		Token: token,
	}, nil
}

func validateLogin(req *ssov1.LoginRequest) error {
	if err := validator.New().Struct(req.Email); err != nil {
		status.Error(codes.InvalidArgument, "email is incorrect")
	}

	if err := validator.New().Struct(req.Password); err != nil {
		status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == emptyValue {
		status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {

	if err := validateRegister(req); err != nil {
		return nil, err
	}

	// Calling the service layer through interface
	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		// TODO... add errors differentiation
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.RegisterResponse{
		UserId: userID,
	}, nil
}

func validateRegister(req *ssov1.RegisterRequest) error {
	if err := validator.New().Struct(req.Email); err != nil {
		status.Error(codes.InvalidArgument, "email is incorrect")
	}

	if err := validator.New().Struct(req.Password); err != nil {
		status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {

	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		//TODO - add errors differentiation
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func validateIsAdmin(req *ssov1.IsAdminRequest) error {

	if req.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	return nil
}
