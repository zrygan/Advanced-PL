package main

import "fmt"

func monthlyIr(ir float32) float32 {
	return ir / 100 / 12 // Convert percentage to decimal and then to monthly
}

func loanTermMonths(loanTerm int32) int32 {
	return loanTerm * 12
}

func totalInterest(loanAmount float64, monthlyIr float32, loanTermMonths int32) float64 {
	return loanAmount * float64(monthlyIr) * float64(loanTermMonths)
}

func monthlyRepayment(loanAmount float64, totalInterest float64, loanTermMonths int32) float64 {
	return (loanAmount + totalInterest) / float64(loanTermMonths)
}

func main() {
	// the amount of the loan
	// float64 for high precision
	var loanAmount float64

	// the interest rate (the percentage part)
	// if interest rate is 3.2% then interest_rate = 3.2
	var interestRate float32

	// the loan term (number of years)
	var loanTerm int32

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
