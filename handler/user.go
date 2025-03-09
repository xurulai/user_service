package handler

import (
	"context"
	"fmt"
	"user_service/biz/user"
	"user_service/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserSrv struct {
	proto.UnimplementedUserServiceServer // 嵌入未实现的 OrderServer 接口，用于兼容 gRPC 的接口实现
}

// 用户注册
// Register 用户注册
func (u *UserSrv) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("in CreateOrder ... ") // 打印进入方法的日志

	// 参数处理
	if req.GetUsername() == "" || req.GetPassword() == "" || req.GetEmail() == "" {
		// 无效的请求
		return nil, status.Error(codes.InvalidArgument, "请求参数有误") // 返回 gRPC 的 InvalidArgument 错误
	}

	// 业务处理
	response, err := user.Register(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		zap.L().Error("user.Register failed", zap.Error(err)) // 记录错误日志
		return response, status.Error(codes.Internal, "内部错误") // 返回 gRPC 的 Internal 错误
	}

	return response, nil // 返回空响应，表示操作成功
}
