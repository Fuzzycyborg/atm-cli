# atm-cli

## Query

A simple ATM cli program that has the following features:

- Change Pin (you can create a default pin in your code and then change it in the function)
- Account balance
- Withdraw funds
- Deposit funds
- Cancel/Exit (selecting this option should exit the program)

The flow of the program should be as follows:

- Run the program
- The program should prompt the user to enter their pin
- If the pin is incorrect, the program should display an error message and prompt the user to enter their pin again
- If the pin is correct, the program should display a menu with the following options:
  - Change Pin
  - Account Balance
  - Withdraw Funds
  - Deposit Funds
  - Cancel/Exit
- The user should be able to select an option from the menu
- If the user selects any other option, the program should display a message that the option is not yet implemented and then display the menu again
- If the user selects the cancel/exit option, the program should exit