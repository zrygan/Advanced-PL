=begin
  Last names: Miranda
  Language: Ruby
  Paradigm(s):Multi-paradigm: Object-Oriented Programming, Functional, and Generic
=end

# no includes and imports necessary
# no need to establish a main function

# Ruby has print(), puts() which works like println, and printf() like in C
# but printf was used to keep the code simple

# no need to establish the datatypes for variables because of dynamic typing

# To get inputs, you may use gets.chomp
# but to ensure uniformity in outputs and avoid them being assigned as strings, .to_f and .to_i is added to explicitly typecast
# automatic typecasting to string despite purely numerical inputs is a problem in Ruby

printf("Loan Amount: ")
loan_amount = gets.chomp.to_f
printf("Annual Interest Rate in Percent: ")
interest_rate = gets.chomp.to_f / 100
monthly_interest_rate = interest_rate / 12
printf("Loan Term in Years: ")
loan_term = gets.chomp.to_i * 12

total_interest = loan_amount * monthly_interest_rate * loan_term
monthly_repayment = (loan_amount + total_interest)/loan_term 

# This is put to avoid edge cases such as all inputs being zero
# Monthly repayment can become NaN since division by 0 can result from the loan_term
if monthly_repayment.nan?
  monthly_repayment = 0
end

printf("\nLoan Amount: PHP %.2f\n", loan_amount)
printf("Annual Interest Rate: %.2f%\n", interest_rate * 100)
printf("Loan Term: %d months\n", loan_term)
printf("Monthly Repayment: PHP %.2f\n", monthly_repayment)
printf("Total Interest: PHP %.2f\n", total_interest)