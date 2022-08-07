package main

import (
	"fmt"

	"github.com/setcreed/hade-ioc/examples/services"
	"github.com/setcreed/hade-ioc/pkg/injector"
)

func main() {
	injector.BeanFactory.Set(services.NewOrderService())
	order := injector.BeanFactory.Get((*services.OrderService)(nil))
	fmt.Println(order)
}
