=begin
  Last names: Miranda, Ching, Ganituen
  Language: Ruby
  Paradigm(s):Multi-paradigm: Object-Oriented Programming, Functional, and Generic
=end

# no includes and imports necessary
# no need to establish a main function

# Ruby has print(), puts() which works like println, and printf() like in C
# no need to establish the datatypes for variables because of dynamic typing

# To get inputs, you may use gets.chomp
# but to ensure uniformity in outputs and avoid them being assigned as strings, .to_f and .to_i is added to explicitly typecast
# automatic typecasting to string despite purely numerical inputs is a problem in Ruby

print("Loan Amount: ")
loan_amount = gets.chomp.to_f
print("Annual Interest Rate in Percent: ")
interest_rate = gets.chomp.to_f / 100
monthly_interest_rate = interest_rate / 12
print("Loan Term in Years: ")
loan_term = gets.chomp.to_i * 12

total_interest = loan_amount * monthly_interest_rate * loan_term
monthly_repayment = (loan_amount + total_interest)/loan_term

printf("\nLoan Amount: PHP %.2f\n", loan_amount)
puts("Annual Interest Rate: #{interest_rate}%")
puts("Loan Term: #{loan_term} months")
printf("Monthly Repayment: PHP %.2f\n", monthly_repayment)
printf("Total Interest: PHP %.2f\n", total_interest)