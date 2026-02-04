package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

/*
- Given a persons information, determine if they have full access, access to sensitive files, and transactions.
- Information includes if they are an employee, if they are a manager, if they are a customer and account balance.
	- If employee -> full access
	- If manager -> access to sensitive files
	- If customer AND account balance is more than 1000 -> transactions allowed
	- If none, access is denied
*/

func main() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\t-----BANK SECURITY SYSTEM-----")
	fmt.Println(strings.Repeat("-", 50))

	userSecurityCheck := securityCheck()

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\t-----USER REPORT-----")
	fmt.Printf("FULL ACCESS: %t\n", userSecurityCheck[0])
	fmt.Printf("SENSITIVE FILE ACCESS: %t\n", userSecurityCheck[1])
	fmt.Printf("ABLE TO TRANSACT: %t\n", userSecurityCheck[2])
	fmt.Println(strings.Repeat("-", 50))
}

func getInfo() []bool {

	checkAnswer := func(answer int) bool {
		return answer == 1
	}

	fmt.Println("\t---ENTER 1 FOR YES---")
	fmt.Println("\t---ENTER 2 FOR NO---")

	userAnswer := GetValidatedNumber("Are you an employee: ", 1, 2)
	isEmployee := checkAnswer(userAnswer)

	userAnswer = GetValidatedNumber("Are you a manager: ", 1, 2)
	isManager := checkAnswer(userAnswer)

	userAnswer = GetValidatedNumber("Are you a customer: ", 1, 2)
	isCustomer := checkAnswer(userAnswer)

	userAnswer = GetValidatedNumber("Enter your balance: ", 0, 999999)
	goodBalance := false
	if userAnswer > 1000 {
		goodBalance = true
	}

	return []bool{isEmployee, isManager, isCustomer, goodBalance}
}

func securityCheck() []bool {
	userInfo := getInfo()
	fmt.Println(userInfo)

	fullAccess := false
	sensFileAccess := false
	transactionAccess := false

	if userInfo[0] {
		fullAccess = true
	}
	if userInfo[1] {
		sensFileAccess = true
	}
	if userInfo[2] && userInfo[3] {
		transactionAccess = true
	}

	return []bool{fullAccess, sensFileAccess, transactionAccess}
}


// Helper func to prompt user for integer type response, return that integer
func GetValidatedNumber(prompt string, min, max int) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			fmt.Printf("Number must be between %d and %d. Try again.\n", min, max)
		} else {
			fmt.Println("Invalid number. Please try again.")
		}
	}
}
