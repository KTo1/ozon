// adapters/driving/http/handler.go
package http

import (
	"encoding/json"
	driving "example/ports/driving"
	"net/http"
)

type OrderHandler struct {
	service driving.OrderService
}

func NewOrderHandler(svc driving.OrderService) *OrderHandler {
	return &OrderHandler{service: svc}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
		Amount int    `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	order, err := h.service.CreateOrder(req.UserID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID string `json:"order_id"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	order, err := h.service.GetOrder(req.OrderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
