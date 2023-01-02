package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status" // 不能引用 internal的status包，会报错
)

const (
	ClientIdKey     = "client-id"
	ClientSecretKey = "client-secret"
)

// 通过暴露Auth函数，来提供grpc中间件函数
func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return newGrpcAuther().Auth
}

// 通过暴露StreamAuth函数，来提供grpc stream中间件函数
func GrpcAuthStreamServerInterceptor() grpc.StreamServerInterceptor {
	return newGrpcAuther().StreamAuth
}

type grpcAuther struct {
}

func newGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

// 认证逻辑
// http/1.1中有header，grpc使用http/2.0，在grpc中叫做metadata，metadata等价于header
func (a *grpcAuther) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 从ctx中获取到了metadata信息，所以中间信息都会通过metadata传递，服务端 <-- 客户端
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md)
	}

	// 认证请求合法性
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	err = a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	// 认证通过，请求交给后面的handler处理
	return handler(ctx, req)
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "client credential %s, %s not provided", ClientIdKey, ClientSecretKey)
	}

	if !(clientId == "admin" && clientSecret == "123456") {
		return status.Errorf(codes.Unauthenticated, "client-id or client-secret is not correct")
	}

	return nil
}

// 从metadata中获取客户端凭证
func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cids := md.Get(ClientIdKey)
	sids := md.Get(ClientSecretKey)

	if len(cids) > 0 {
		clientId = cids[0]
	}

	if len(sids) > 0 {
		clientSecret = sids[0]
	}
	return
}

func (a *grpcAuther) StreamAuth(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	md, _ := metadata.FromIncomingContext(ss.Context())

	// 认证请求合法性
	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)
	err := a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return err
	}

	// 认证通过，请求交给后面的handler处理
	return handler(srv, ss)
}
