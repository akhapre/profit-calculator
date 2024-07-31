package main

import "fmt"

func main() {

	var revenue float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	var expenses float64
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	var taxRate float64
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	fmt.Printf("Earning before tax : %.2f\n", earningsBeforeTax)

	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	fmt.Printf("Earning after tax : %.2f\n", earningsAfterTax)

	ratio := earningsBeforeTax / earningsAfterTax
	fmt.Printf("Ratio: %.4f\n", ratio)

}
