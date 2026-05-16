package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Key type to avoid collisions
type key string

const traceKey key = "traceID"

// Response struct for JSON
type OrderResponse struct {
	TraceID     string `json:"trace_id"`
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	StockStatus string `json:"stock_status"`
	Error       string `json:"error,omitempty"`
}

// Fake DB query
func queryDB(ctx context.Context, orderID string) (string, error) {
	select {
	case <-time.After(600 * time.Millisecond): // 600ms lag
		return "shipped", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Fake RPC call
func callStockService(ctx context.Context, orderID string) (string, error) {
	select {
	case <-time.After(400 * time.Millisecond): // 400ms lag
		return "in_stock", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Main handler
func handleOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	if orderID == "" {
		http.Error(w, "Need an order_id, buddy!", http.StatusBadRequest)
		return
	}

	// 1.5s timeout + client cancel support
	ctx, cancel := context.WithTimeout(r.Context(), 1000*time.Millisecond)
	defer cancel()

	// Toss in a trace ID
	traceID := fmt.Sprintf("req-%d", time.Now().UnixNano())
	ctx = context.WithValue(ctx, traceKey, traceID)

	// Early exit if client bails
	select {
	case <-ctx.Done():
		resp := OrderResponse{TraceID: traceID, Error: "Client ditched"}
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode(resp)
		log.Printf("Canceled [traceID=%s]: %v", traceID, ctx.Err())
		return
	default:
	}

	// Run DB and RPC in one goroutine
	type result struct {
		status string
		stock  string
		err    error
	}
	ch := make(chan result, 1)

	go func() {
		status, err := queryDB(ctx, orderID)
		if err != nil {
			ch <- result{err: err}
			return
		}
		stock, err := callStockService(ctx, orderID)
		ch <- result{status: status, stock: stock, err: err}
	}()

	// Wait for results or timeout
	var resp OrderResponse
	select {
	case res := <-ch:
		if res.err != nil {
			resp = OrderResponse{TraceID: traceID, OrderID: orderID, Error: res.err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)
			log.Printf("Failed [traceID=%s]: %v", traceID, res.err)
			return
		}
		resp = OrderResponse{TraceID: traceID, OrderID: orderID, Status: res.status, StockStatus: res.stock}
	case <-ctx.Done():
		resp = OrderResponse{TraceID: traceID, OrderID: orderID, Error: "Timed out"}
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(resp)
		log.Printf("Timeout [traceID=%s]: %v", traceID, ctx.Err())
		return
	}

	// Happy path
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	log.Printf("Done [traceID=%s]", traceID)
}

func main() {
	http.HandleFunc("/api/order", handleOrder)
	log.Println("Firing up on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
