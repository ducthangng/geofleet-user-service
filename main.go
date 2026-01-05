package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/ducthangng/geofleet-proto/user"
	"github.com/ducthangng/geofleet/user-service/registry"
	"github.com/ducthangng/geofleet/user-service/singleton"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var devenv string

func main() {

	// gather the configuration of the service
	cfg := singleton.ReadConfig(devenv)

	// start a global context
	ctx := context.Background()

	// connect to postgre database
	singleton.ConnectPostgre(ctx)

	// connect to redis
	// singleton.ConnectRedis()

	// 1. Mở port TCP
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.ServerHost, cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Khởi tạo gRPC Server
	s := grpc.NewServer()

	userHandler, err := registry.Initialize(ctx)
	if err != nil {
		log.Panic(err)
		return
	}

	pb.RegisterUserServiceServer(s, userHandler)

	// 3. Đăng ký với Consul
	err = singleton.RegisterWithConsul(cfg.Server.ServiceID, cfg.Server.ServiceName, cfg.Server.Port)
	if err != nil {
		log.Fatalf("Lỗi đăng ký Consul: %v", err)
	}

	// 1. Khởi tạo Health Server
	healthServer := health.NewServer()

	// 2. Đăng ký Health Service vào gRPC Server của bạn
	healthpb.RegisterHealthServer(s, healthServer)

	// 3. Đặt trạng thái là SERVING (Đang hoạt động)
	// "user-service" ở đây phải khớp với Service Name bạn đăng ký trên Consul
	healthServer.SetServingStatus("user-service", healthpb.HealthCheckResponse_SERVING)

	log.Printf("User Service đã đăng ký với Consul tại port %d", cfg.Server.ConsulPort)

	// 4. Xử lý Graceful Shutdown (Hủy đăng ký khi tắt app)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("Đang tắt Service...")
		// Hủy đăng ký trên Consul để Gateway không gọi vào nữa
		// (Thực hiện gọi client.Agent().ServiceDeregister(serviceID))
		s.GracefulStop()
		os.Exit(0)
	}()

	// 5. Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
