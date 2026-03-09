package service

import (
	"encoding/json"
	"order-system/internal/kafka"
	"order-system/internal/model"
	"order-system/internal/repository"
)

type OrderService struct {
	repo     *repository.OrderRepository
	producer *kafka.Producer
}

func NewOrderService(repo *repository.OrderRepository, producer *kafka.Producer) *OrderService {
	return &OrderService{repo: repo, producer: producer}
}
func (s *OrderService) Create(order model.Order) error {
	err := s.repo.Create(order)
	if err != nil {
		return err
	}
	data, _ := json.Marshal(order)
	return s.producer.Send(data)
}
