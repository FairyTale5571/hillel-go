package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

func CalculateFinalPrice(p Product) (float64, error) {
	if p.Price < 0 || p.Discount < 0 || p.Discount > 100 {
		return 0, fmt.Errorf("invalid input: price and discount must be between 0 and 100")
	}
	result := (p.Price * (100 - p.Discount)) / 100
	return result, nil
}

func DisplayProductDetails(p Product) {
	finalPrice, err := CalculateFinalPrice(p)
	if err != nil {
		fmt.Printf("Product: %s - Invalid data: %v\n", p.Name, err)
		return
	}
	discountAmount := (p.Price * p.Discount) / 100
	fmt.Printf("Product: %s, Original Price: $%.2f, Discount: %.2f%%, Discount Amount: $%.2f, Final Price: $%.2f\n",
		p.Name, p.Price, p.Discount, discountAmount, finalPrice)
}

func CalculateTotals(products []Product) (float64, float64, error) {
	var totalOriginal, totalFinal float64
	for _, product := range products {
		finalPrice, err := CalculateFinalPrice(product)
		if err != nil {
			return 0, 0, err
		}
		totalOriginal += product.Price
		totalFinal += finalPrice
	}
	return totalOriginal, totalFinal, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var products []Product

	for {
		fmt.Println("Enter product name (or type 'done' to finish):")
		scanner.Scan()
		name := scanner.Text()
		if strings.ToLower(name) == "done" {
			break
		}

		fmt.Println("Enter price:")
		scanner.Scan()
		priceInput := scanner.Text()
		price, err := strconv.ParseFloat(priceInput, 64)
		if err != nil || price < 0 {
			fmt.Println("Invalid price. Please enter a non-negative number.")
			continue
		}

		fmt.Println("Enter discount percentage (0-100%):")
		scanner.Scan()
		discountInput := scanner.Text()
		discount, err := strconv.ParseFloat(discountInput, 64)
		if err != nil || discount < 0 || discount > 100 {
			fmt.Println("Invalid discount. Discount must be between 0% and 100%.")
			continue
		}

		product := Product{Name: name, Price: price, Discount: discount}
		products = append(products, product)
		DisplayProductDetails(product)
	}

	totalOriginal, totalFinal, err := CalculateTotals(products)
	if err != nil {
		fmt.Println("Error calculating totals:", err)
		return
	}
	fmt.Printf("Total Original Price: $%.2f\n", totalOriginal)
	fmt.Printf("Total Final Price after Discounts: $%.2f\n", totalFinal)
}
