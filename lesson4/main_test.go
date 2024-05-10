package main

import (
	"testing"
)

func TestCalculateFinalPrice_ValidInput(t *testing.T) {
	product := Product{Name: "Name1", Price: 100.0, Discount: 20.0}
	expected := 80.0
	result, err := CalculateFinalPrice(product)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expected {
		t.Errorf("TestCalculateFinalPrice failed: expected %f, got %f", expected, *result)
	}
}

func TestCalculateFinalPrice_NegativePrice(t *testing.T) {
	product := Product{Name: "Name1", Price: -100.0, Discount: 10.0}
	_, err := CalculateFinalPrice(product)
	if err == nil {
		t.Errorf("Expected error for negative price, got nil")
	}
}

func TestCalculateFinalPrice_DiscountOver100(t *testing.T) {
	product := Product{Name: "Name1", Price: 100.0, Discount: 150.0}
	_, err := CalculateFinalPrice(product)
	if err == nil {
		t.Errorf("Expected error for discount over 100%%, got nil")
	}
}

func TestCalculateFinalPrice_NegativeDiscount(t *testing.T) {
	product := Product{Name: "Name1", Price: 100.0, Discount: -20.0}
	_, err := CalculateFinalPrice(product)
	if err == nil {
		t.Errorf("Expected error for negative discount, got nil")
	}
}

func TestCalculateTotals_NoDiscounts(t *testing.T) {
	products := []Product{{Name: "Name1", Price: 200.0, Discount: 0.0}}
	expectedOriginal := 200.0
	expectedFinal := 200.0
	original, final, err := CalculateTotals(products)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *original != expectedOriginal || *final != expectedFinal {
		t.Errorf("TestCalculateTotals_NoDiscounts failed: expected %f and %f, got %f and %f", expectedOriginal, expectedFinal, *original, *final)
	}
}

func TestCalculateTotals_WithDiscounts(t *testing.T) {
	products := []Product{{Name: "Name1", Price: 200.0, Discount: 25.0}}
	expectedOriginal := 200.0
	expectedFinal := 150.0
	original, final, err := CalculateTotals(products)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *original != expectedOriginal || *final != expectedFinal {
		t.Errorf("TestCalculateTotals_WithDiscounts failed: expected %f and %f, got %f and %f", expectedOriginal, expectedFinal, *original, *final)
	}
}

func TestMultipleProducts(t *testing.T) {
	products := []Product{
		{Name: "Name1", Price: 100.0, Discount: 10.0},
		{Name: "Name2", Price: 200.0, Discount: 20.0},
	}
	_, _, err := CalculateTotals(products)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(products) != 2 {
		t.Errorf("Expected 2 products, got %d", len(products))
	}
}

func TestZeroDiscount(t *testing.T) {
	product := Product{Name: "Name1", Price: 100.0, Discount: 0.0}
	expected := 100.0
	result, err := CalculateFinalPrice(product)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expected {
		t.Errorf("TestZeroDiscount failed: expected %f, got %f", expected, *result)
	}
}

func TestZeroPrice(t *testing.T) {
	product := Product{Name: "Freebie", Price: 0.0, Discount: 20.0}
	expected := 0.0
	result, err := CalculateFinalPrice(product)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expected {
		t.Errorf("TestZeroPrice failed: expected %f, got %f", expected, *result)
	}
}

func TestFinalSums(t *testing.T) {
	products := []Product{
		{Name: "Name1", Price: 100.0, Discount: 10.0},
		{Name: "Name2", Price: 200.0, Discount: 20.0},
	}
	expectedOriginal := 300.0
	expectedFinal := 250.0
	original, final, err := CalculateTotals(products)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *original != expectedOriginal || *final != expectedFinal {
		t.Errorf("TestFinalSums failed: expected %f and %f, got %f and %f", expectedOriginal, expectedFinal, *original, *final)
	}
}
