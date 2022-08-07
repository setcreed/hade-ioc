package services

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0"}
}

func (o *OrderService) GetOrderInfo(uid int) {
	fmt.Printf("获取用户id=%d的订单信息", uid)
}
