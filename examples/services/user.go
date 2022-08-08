package services

import "fmt"

type UserService struct {
	Order *OrderService `inject:"ServiceConfig.OrderService()"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserInfo(uid int) {
	fmt.Printf("获取用户id=%d的详细信息", uid)
}
