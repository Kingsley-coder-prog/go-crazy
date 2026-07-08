package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your account balance is: $", accountBalance)
		case 2:
			fmt.Print("Your deposit: $")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount: $", accountBalance)
		case 3:
			fmt.Print("Your withdraw: $")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds. Your account balance is: $", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
			fmt.Println("Balance updated! New amount: $", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			// break
			return
		}

	}

}
