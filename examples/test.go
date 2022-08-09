package main

import (
	"fmt"

	"github.com/setcreed/hade-ioc/examples/config"
	"github.com/setcreed/hade-ioc/examples/services"
	"github.com/setcreed/hade-ioc/pkg/injector"
)

func main() {
	//injector.BeanFactory.Set(services.NewOrderService())
	//userService := services.NewUserService()
	//injector.BeanFactory.Apply(userService)
	//
	//fmt.Println(userService.Order)

	serviceConfig := config.NewServiceConfig()
	//injector.BeanFactory.ExprMap = map[string]interface{}{
	//	"ServiceConfig": serviceConfig,
	//}
	//order := services.NewOrderService()
	//order.Version = "2.0"
	injector.BeanFactory.Config(serviceConfig)

	{
		userService := services.NewUserService()
		injector.BeanFactory.Apply(userService)
		fmt.Println(userService.Order)
	}

	{
		adminService := services.NewAdminService()
		injector.BeanFactory.Apply(adminService)
		fmt.Println(adminService.Order)
	}

}
