import java.util.Scanner

fun main() {
    val scanner = Scanner(System.`in`)

    print("Loan Amount: PHP ")
    val loan = scanner.nextDouble()

    print("Annual Interest Rate: ")
    val rate = scanner.nextDouble()

    val ratePercent = rate / 100

    print("Loan Term: ")
    val duration = scanner.nextInt()

    val monthlyRate = ratePercent / 12
    val loanMonths = duration * 12

    val totalInterest = loan * monthlyRate * loanMonths

    val repayment = (loan + totalInterest) / loanMonths

    println("\nLoan Amount: PHP $loan")
    println("Annual Interest Rate: $rate%")
    println("Loan Term: $loanMonths months")
    println("Monthly Repayment: PHP ${"%,.2f".format(repayment)}")
    println("Total Interest: PHP ${"%,.0f".format(totalInterest)}")
}