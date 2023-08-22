package service

import (
	"context"

	"github.com/aclgo/grpc-jwt/proto"
)

func (us *UserService) Register(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreatedUserResponse, error) {
	return nil, nil
}
func (us *UserService) Login(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	return nil, nil
}
func (us *UserService) Logout(ctx context.Context, req *proto.UserLogoutRequest) (*proto.UserLogoutResponse, error) {
	return nil, nil
}
func (us *UserService) FindById(ctx context.Context, req *proto.FindByIdRequest) (*proto.FindByIdResponse, error) {
	return nil, nil
}
func (us *UserService) FindByEmail(ctx context.Context, req *proto.FindByEmailRequest) (*proto.FindByEmailResponse, error) {
	return nil, nil
}
