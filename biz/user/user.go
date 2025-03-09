package user

import (
	"context"
	"user_service/dao/mysql"
	"user_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register 用户注册
func Register(ctx context.Context, username string, password string, email string) (*proto.RegisterResponse, error) {
	// 调用 mysql 包中的 RegisterUser 方法
	err := mysql.RegisterUser(ctx, username, password, email)
	if err != nil {
		// 如果注册失败，返回错误
		return &proto.RegisterResponse{
			Success: false,
			Message: "注册用户失败",
		}, status.Error(codes.Internal, "内部错误")
	}

	// 注册成功，返回成功响应
	return &proto.RegisterResponse{
		Success: true,
		Message: "注册用户成功",
	}, nil
}
