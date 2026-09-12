package main

import (
	"errors"
	"fmt"
)

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Voucher struct {
	Code            string
	DiscountPercent float64
	MaxDiscount     float64
	MinPurchase     float64
}

func CalculateFinalPrice(items []CartItem, voucher *Voucher) (subtotal float64, discount float64, total float64, err error) {
	if len(items) == 0 {
		return 0, 0, 0, errors.New("cart cannot be empty")
	}

	for _, item := range items {
		if item.Quantity <= 0 || item.Price < 0 {
			return 0, 0, 0, errors.New("invalid item price or quantity")
		}
		subtotal += item.Price * float64(item.Quantity)
	}

	if voucher == nil {
		return subtotal, 0, subtotal, nil
	}

	if subtotal < voucher.MinPurchase {
		return subtotal, 0, subtotal, nil
	}
	discount = subtotal * (voucher.DiscountPercent / 100)

	if voucher.MaxDiscount > 0 && discount > voucher.MaxDiscount {
		discount = voucher.MaxDiscount
	}

	total = subtotal - discount
	return subtotal, discount, total, nil
}

func main() {
	promoVoucher := &Voucher{
		Code:            "PROMO10",
		DiscountPercent: 10.0,
		MaxDiscount:     50000.0,
		MinPurchase:     200000.0,
	}

	// Test case 1: Cart with a Max Capped Discount
	cart1 := []CartItem{
		{ProductID: "P001", Name: "Mechanical Keyboard", Price: 750000, Quantity: 1},
	}
	subtotal, discount, total, err := CalculateFinalPrice(cart1, promoVoucher)
	if err != nil {
		panic(err)
	}
	fmt.Printf("===== Test Case 1 =====\n")
	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Discount: %.2f\n", discount)
	fmt.Printf("Total: %.2f\n", total)

	// Test case 2: Cart with a Min Purchase Requirement Not Met
	cart2 := []CartItem{
		{ProductID: "P002", Name: "Wireless Mouse", Price: 150000, Quantity: 1},
	}
	subtotal, discount, total, err = CalculateFinalPrice(cart2, promoVoucher)
	if err != nil {
		panic(err)
	}
	fmt.Printf("===== Test Case 2 =====\n")
	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Discount: %.2f\n", discount)
	fmt.Printf("Total: %.2f\n", total)

	// Test case 3: Cart with No Voucher Applied
	cart3 := []CartItem{
		{ProductID: "P003", Name: "Gaming Headset", Price: 300000, Quantity: 1},
	}
	subtotal, discount, total, err = CalculateFinalPrice(cart3, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("===== Test Case 3 =====\n")
	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Discount: %.2f\n", discount)
	fmt.Printf("Total: %.2f\n", total)

	// Test case 4: Cart with Invalid Item Price or Quantity
	cart4 := []CartItem{
		{ProductID: "P004", Name: "USB-C Hub", Price: -50000, Quantity: 1},
	}
	subtotal, discount, total, err = CalculateFinalPrice(cart4, promoVoucher)
	if err != nil {
		fmt.Printf("===== Test Case 4 =====\n")
		fmt.Printf("Error: %v\n", err)
	}
}
