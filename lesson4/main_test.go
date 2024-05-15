package main

import (
	"testing"
)

type CalculateTotalsTest struct {
	name             string
	products         []Product
	expectedOriginal float64
	expectedFinal    float64
	wantErr          bool
}

type CalculateFinalPriceTest struct {
	name     string
	product  Product
	expected float64
	wantErr  bool
}

func TestCalculate(t *testing.T) {
	totalTests := []CalculateTotalsTest{
		{"No Discounts", []Product{{Name: "Name1", Price: 200.0, Discount: 0.0}}, 200.0, 200.0, false},
		{"With Discounts", []Product{{Name: "Name1", Price: 200.0, Discount: 25.0}}, 200.0, 150.0, false},
		{"Multiple Products", []Product{{Name: "Name1", Price: 100.0, Discount: 10.0}, {Name: "Name2", Price: 200.0, Discount: 20.0}}, 300.0, 250.0, false},
	}

	finalPriceTests := []CalculateFinalPriceTest{
		{"Valid Input", Product{Name: "Name1", Price: 100.0, Discount: 20.0}, 80.0, false},
		{"Negative Price", Product{Name: "Name1", Price: -100.0, Discount: 10.0}, 0.0, true},
		{"Discount Over 100", Product{Name: "Name1", Price: 100.0, Discount: 150.0}, 0.0, true},
		{"Negative Discount", Product{Name: "Name1", Price: 100.0, Discount: -20.0}, 0.0, true},
		{"Zero Discount", Product{Name: "Name1", Price: 100.0, Discount: 0.0}, 100.0, false},
		{"Zero Price", Product{Name: "Freebie", Price: 0.0, Discount: 20.0}, 0.0, false},
	}

	// Running CalculateTotals tests
	for _, tt := range totalTests {
		original, final, err := CalculateTotals(tt.products)
		if (err != nil) != tt.wantErr {
			if err != nil {
				t.Errorf("CalculateTotals %s: unexpected error: %v", tt.name, err)
			} else {
				t.Errorf("CalculateTotals %s: expected error but got none", tt.name)
			}
		}
		if original != tt.expectedOriginal || final != tt.expectedFinal {
			t.Errorf("CalculateTotals %s: got %v and %v, want %v and %v", tt.name, original, final, tt.expectedOriginal, tt.expectedFinal)
		}
	}

	// Running CalculateFinalPrice tests
	for _, tt := range finalPriceTests {
		result, err := CalculateFinalPrice(tt.product)
		if (err != nil) != tt.wantErr {
			if err != nil {
				t.Errorf("CalculateFinalPrice %s: unexpected error: %v", tt.name, err)
			} else {
				t.Errorf("CalculateFinalPrice %s: expected error but got none", tt.name)
			}
		}
		if result != tt.expected {
			t.Errorf("CalculateFinalPrice %s: got %v, want %v", tt.name, result, tt.expected)
		}
	}
}
