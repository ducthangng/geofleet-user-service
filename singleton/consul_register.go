package singleton

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
)

func getLocalIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func RegisterWithConsul(serviceID string, serviceName string, port int) error {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" // Địa chỉ Consul Agent

	cfg := GetConfig()

	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	// address := "127.0.0.1" // docker
	// if cfg.Server.Env == "dev" {
	// 	// QUAN TRỌNG: Bảo Consul gọi ngược ra Mac Host
	// 	address = "host.docker.internal"
	// }

	address := getLocalIP()
	log.Println("connecting to address: ", address)

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{
			// Kiểm tra gRPC health check (Consul sẽ gọi vào port này để xem service còn sống không)
			GRPC:                           fmt.Sprintf("%s:%d", address, cfg.Server.Port),
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	return client.Agent().ServiceRegister(registration)
}
