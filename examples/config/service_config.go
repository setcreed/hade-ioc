package config

import "github.com/setcreed/hade-ioc/examples/services"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (s *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}
