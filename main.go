package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// initial account balance
const balance = 1000

// default pin
const pin = 1234

func main() {
	// create a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)

	// prompt the user to enter their pin
	fmt.Println("Please enter your PIN:")
	scanner.Scan()
	inputPin, _ := strconv.Atoi(scanner.Text())

	// check if the entered pin is correct
	if inputPin != pin {
		fmt.Println("Invalid PIN. Exiting program.")
		return
	}

	// if the pin is correct, display the menu of options
	for {
		fmt.Println("\nWelcome to the ATM machine. Please select an option:")
		fmt.Println("1. Change PIN")
		fmt.Println("2. Account balance")
		fmt.Println("3. Withdraw funds")
		fmt.Println("4. Deposit funds")
		fmt.Println("5. Cancel/Exit")
		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			// change the pin
			fmt.Println("Please enter your new PIN:")
			scanner.Scan()
			newPin, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Your PIN has been changed to", newPin)
		case 2:
			// display the account balance
			fmt.Println("Your account balance is", balance)
		case 3:
			// withdraw funds
			fmt.Println("Enter the amount to withdraw:")
			scanner.Scan()
			amount, _ := strconv.Atoi(scanner.Text())

			// check if the entered amount is valid
			if amount > balance {
				fmt.Println("Insufficient funds.")
			} else {
				fmt.Println("Please collect your cash.")
			}
		case 4:
			// deposit funds
			fmt.Println("Enter the amount to deposit:")
			scanner.Scan()
			amount, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Your account balance is", balance+amount)
		case 5:
			// exit the program
			fmt.Println("Exiting the program.")
			return
		default:
			// if the entered option is not valid, display an error message
			fmt.Println("Invalid option. Please try again.")
		}

		// ask the user if they want to perform another operation
		fmt.Println("\nDo you want to perform another operation? (yes/no)")
		scanner.Scan()
		response := strings.ToLower(scanner.Text())

		if response != "yes" {
			fmt.Println("Exiting the program.")
			return
		}
	}
}
