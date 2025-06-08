package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Item represents a single purchased item from the receipt.
type Item struct {
	Description string  `json:"shortDescription"`
	Price       float64 `json:"-"`     // Parsed numeric value
	RawPrice    string  `json:"price"` // Original string value from JSON
}

// Receipt represents the receipt JSON payload.
type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"-"`     // Parsed numeric value
	RawTotal     string  `json:"total"` // Original string value from JSON
}

func (r *Receipt) UnmarshalJSON(data []byte) error {
	// Alias to avoid infinite recursion
	type Alias Receipt
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	// Decode raw JSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse total as float
	total, err := strconv.ParseFloat(r.RawTotal, 64)
	if err != nil {
		return fmt.Errorf("invalid total value: %v", err)
	}
	r.Total = total

	// Parse each item's price
	for i := range r.Items {
		price, err := strconv.ParseFloat(r.Items[i].RawPrice, 64)
		if err != nil {
			return fmt.Errorf("invalid item price at index %d: %v", i, err)
		}
		r.Items[i].Price = price
	}

	return nil
}
