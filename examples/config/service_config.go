package config

import (
	"github.com/setcreed/hade-ioc/examples/services"
	"log"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (s *ServiceConfig) OrderService() *services.OrderService {
	log.Println("初始化 orderService")
	return services.NewOrderService()
}
