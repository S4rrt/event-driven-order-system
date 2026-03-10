package handler

import (
	"encoding/json"
	"event-driven-order-system/internal/model"
	"event-driven-order-system/internal/service"
	"net/http"
)

type OrderService struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderService {
	return &OrderService{service: service}
}
func (h *OrderService) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
