package handler

import (
	"encoding/json"
	"net/http"
	"pharmacy-shop/internal/domain"
	"pharmacy-shop/internal/repository"
)

type OrderHandler struct {
	orderRepo *repository.OrderRepository
}

func NewOrderHandler(orderRepo *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepo: orderRepo}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.orderRepo.CreateOrderWithTransaction(&order); err != nil {
		if err == domain.ErrInsufficientStock {
			http.Error(w, "Insufficient stock", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderRepo.GetAllOrders()
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
