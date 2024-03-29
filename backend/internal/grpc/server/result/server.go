package result

import (
	"backend/conf/settings"
	pb "backend/internal/grpc/server/result/pb"
	"backend/pkg/mw"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/sunflower10086/Cococola/etcd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Init() {
	resultConf := settings.Conf.AllServer.ResultConfig
	Addr := fmt.Sprintf("%s:%s", resultConf.Host, resultConf.Port)

	fmt.Println(Addr)
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		mw.SugarLogger.Errorf("net.Listen err: %v", err)
	}

	// 添加中间间
	var opts []grpc.ServerOption
	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_opentracing.StreamServerInterceptor(),
		//grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.StreamServerInterceptor(mw.ZapInterceptor(settings.Conf)),
		//grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
		grpc_recovery.StreamServerInterceptor(mw.RecoveryInterceptor()),
	)))

	// 新建grpc服务器实例
	grpcServer := grpc.NewServer(opts...)
	// 在gRPC服务器注册我们的服务
	pb.RegisterResultServer(grpcServer, &ResultServerImpl{})

	// 在etcd中注册服务
	svc, err := etcd.NewServiceRegister(
		[]string{settings.Conf.EtcdConf.Endpoint},
		"/gRPC/"+resultConf.Name,
		resultConf.GetAddr(),
		3,
	)
	if err != nil {
		mw.SugarLogger.Errorf("etcd Register err: %v", err)
		return
	}
	go svc.ListenLeaseRespChan()

	zap.L().Debug(Addr + " net.Listing...")

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		mw.SugarLogger.Errorf("grpcServer.Serve err: %v", err)
	}
}
