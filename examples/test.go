package main

import (
	"fmt"
	"github.com/setcreed/hade-ioc/examples/services"
	"github.com/setcreed/hade-ioc/pkg/injector"
)

func main() {
	injector.BeanFactory.Set(services.NewOrderService())
	userService := services.NewUserService()
	injector.BeanFactory.Apply(userService)

	fmt.Println(userService.Order)
}
