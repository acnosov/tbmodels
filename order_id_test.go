package tbmodels

import (
	"encoding/json"
	"testing"
)

func TestOrderIDAboveInt32Range_ReproTRA27(t *testing.T) {
	var order OrderDataMessage
	if err := json.Unmarshal([]byte(`{"order_id":2149501836}`), &order); err != nil {
		t.Fatalf("unmarshal OrderDataMessage: %v", err)
	}
	if order.OrderID != 2149501836 {
		t.Fatalf("got %d", order.OrderID)
	}
	var bet BetMessage
	if err := json.Unmarshal([]byte(`{"order_id":2149501836}`), &bet); err != nil {
		t.Fatalf("unmarshal BetMessage: %v", err)
	}
	if bet.OrderID != 2149501836 {
		t.Fatalf("got %d", bet.OrderID)
	}
	var side Side
	if err := json.Unmarshal([]byte(`{"OrderID":2149501836}`), &side); err != nil {
		t.Fatalf("unmarshal Side: %v", err)
	}
	if side.OrderID != 2149501836 {
		t.Fatalf("got %d", side.OrderID)
	}
}
