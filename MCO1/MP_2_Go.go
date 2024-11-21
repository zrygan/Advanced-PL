/*
******************
Last names: Ganituen
Language: Go
Paradigm(s): Imperative, Procedural
******************
*/

package main

import "fmt"

// Function to calculate the monthly interest rate
func monthlyIr(ir float64) float64 {
	return ir / 100 / 12 // Convert percentage to decimal and then to monthly
}

// Function to convert loan term from years to months
func loanTermMonths(loanTerm int64) int64 {
	return loanTerm * 12
}

// Function to calculate the total interest over the loan term
func totalInterest(loanAmount float64, monthlyIr float64, loanTermMonths int64) float64 {
	return loanAmount * monthlyIr * float64(loanTermMonths)
}

// Function to calculate the monthly repayment amount
func monthlyRepayment(loanAmount float64, totalInterest float64, loanTermMonths int64) float64 {
	return (loanAmount + totalInterest) / float64(loanTermMonths)
}

func main() {
	// the amount of the loan
	// float64 for high precision
	var loanAmount float64

	// the interest rate (the percentage part)
	// if interest rate is 3.2% then interest_rate = 3.2
	var interestRate float64

	// the loan term (number of years)
	var loanTerm int64

	fmt.Print("Loan Amount: ")
	fmt.Scan(&loanAmount)
	fmt.Print("Annual Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Print("Loan Term: ")
	fmt.Scan(&loanTerm)

	monthlyInterestRate := monthlyIr(interestRate)
	loanTermMonths := loanTermMonths(loanTerm)
	totalInterest := totalInterest(loanAmount, monthlyInterestRate, loanTermMonths)
	monthlyRepayment := monthlyRepayment(loanAmount, totalInterest, loanTermMonths)

	fmt.Printf("\nLoan Amount: PHP %.2f\n", loanAmount)
	fmt.Printf("Annual Interest Rate: %.2f%%\n", interestRate)
	fmt.Printf("Loan Term: %d months\n", loanTermMonths)
	fmt.Printf("Monthly Repayment: PHP %.2f\n", monthlyRepayment)
	fmt.Printf("Total Interest: PHP %.2f\n", totalInterest)
}
