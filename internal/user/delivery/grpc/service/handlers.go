package service

import (
	"context"

	"github.com/aclgo/grpc-jwt/internal/models"
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/pkg/grpc_errors"
	"github.com/aclgo/grpc-jwt/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (us *UserService) mustEmbedUnimplementedUserServiceServer() {}

func (us *UserService) Register(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreatedUserResponse, error) {
	params := user.ParamsCreateUser{
		Name:     req.Name,
		Lastname: req.LastName,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
	}

	created, err := us.userUC.Register(ctx, &params)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrors(err), "useUC.Register: %v", err)
	}

	return &proto.CreatedUserResponse{User: parseModelsToProto(created)}, nil
}
func (us *UserService) Login(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	email, password := req.Email, req.Password
	if email == "" || password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Login: %v", grpc_errors.EmptyCredentials{})
	}

	tokens, err := us.userUC.Login(ctx, email, password)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrors(err), "Login: %v", err)
	}

	return &proto.UserLoginResponse{
		Tokens: &proto.Tokens{
			AccessToken:  tokens.Access,
			RefreshToken: tokens.Refresh,
		},
	}, nil
}
func (us *UserService) Logout(ctx context.Context, req *proto.UserLogoutRequest) (*proto.UserLogoutResponse, error) {
	accessTK, refreshTK := req.AccessToken, req.RefreshToken

	if err := us.userUC.Logout(ctx, accessTK, refreshTK); err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrors(err), "Logout: %v", err)
	}

	return &proto.UserLogoutResponse{}, nil
}
func (us *UserService) FindById(ctx context.Context, req *proto.FindByIdRequest) (*proto.FindByIdResponse, error) {
	id := req.Id

	found, err := us.userUC.FindByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrors(err), "userUC.FindByID: %vs", err)
	}

	return &proto.FindByIdResponse{
		User: parseModelsToProto(found),
	}, nil
}
func (us *UserService) FindByEmail(ctx context.Context, req *proto.FindByEmailRequest) (*proto.FindByEmailResponse, error) {
	email := req.Email

	found, err := us.userUC.FindByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrors(err), "userUC.FindByEmail: %v", err)
	}

	return &proto.FindByEmailResponse{
		User: parseModelsToProto(found),
	}, nil
}

func (us *UserService) parseProtoToModels(req *proto.User) *models.User {
	return &models.User{
		Id:        req.Id,
		Name:      req.Name,
		Lastname:  req.LastName,
		Password:  req.Password,
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: req.CreatedAt.AsTime(),
		UpdatedAt: req.UpdatedAt.AsTime(),
	}
}

func parseModelsToProto(user *user.ParamsOutputUser) *proto.User {
	return &proto.User{
		Id:        user.Id,
		Name:      user.Name,
		LastName:  user.Lastname,
		Password:  user.Password,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
