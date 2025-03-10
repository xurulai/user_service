package user

import (
	"context"
	"user_service/dao/mysql"
	"user_service/proto"
	"user_service/third_party/jwt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register 用户注册
func Register(ctx context.Context, username string, password string, email string) (*proto.RegisterResponse, error) {
	// 调用 mysql 包中的 RegisterUser 方法
	// 对密码进行加密
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return &proto.RegisterResponse{
			Success: false,
			Message: "密码加密失败",
		}, status.Error(codes.Internal, "内部错误")
	}
	err = mysql.RegisterUser(ctx, username, hashedPassword, email)
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

// Login 用户登录
func Login(ctx context.Context, username string, password string) (*proto.LoginResponse, error) {
	// 调用 mysql 包中的 GetUserByUsername 方法获取用户信息
	user, err := mysql.GetUserByUsername(ctx, username)
	if err != nil {
		// 如果用户不存在或获取用户信息失败，返回错误
		return &proto.LoginResponse{
			Success: false,
			Message: "用户名或密码错误",
		}, status.Error(codes.Unauthenticated, "用户名或密码错误")
	}

	// 验证密码是否正确
	if err := CheckPassword(user.Password, password); err != nil {
		return &proto.LoginResponse{
			Success: false,
			Message: "用户名或密码错误",
		}, status.Error(codes.Unauthenticated, "用户名或密码错误")
	}

	// 密码正确，生成 JWT Token
	tokenString, err := jwt.GenToken(user.ID, user.Username)
	if err != nil {
		return &proto.LoginResponse{
			Success: false,
			Message: "生成Token失败",
		}, status.Error(codes.Internal, "内部错误")
	}

	// 返回登录成功响应
	return &proto.LoginResponse{
		Success: true,
		Message: "登录成功",
		Token:   tokenString,
	}, nil
}
