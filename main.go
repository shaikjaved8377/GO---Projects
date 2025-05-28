package main

import "fmt"

// Structure to represent a bank account
type BankAccount struct {
	AccountHolder    string
	AccountNumber    int
	Balance          int
	Transactions     [5]int // Array to store the last 5 transactions
	AllTransactions  []int  // Slice to store all transactions
	TransactionIndex int    // Index for circular array
}

// Method to deposit money
func (account *BankAccount) Deposit(amount int) {
	if amount > 0 {
		account.Balance += amount
		account.recordTransaction(amount)
		fmt.Println("Deposit successful! New balance:", account.Balance)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}

// Method to withdraw money
func (account *BankAccount) Withdraw(amount int) {
	if amount > 0 && amount <= account.Balance {
		account.Balance -= amount
		account.recordTransaction(-amount)
		fmt.Println("Withdrawal successful! New balance:", account.Balance)
	} else if amount > account.Balance {
		fmt.Println("Insufficient balance!")
	} else {
		fmt.Println("Invalid withdrawal amount.")
	}
}

// Method to record a transaction
func (account *BankAccount) recordTransaction(amount int) {
	account.Transactions[account.TransactionIndex%5] = amount
	account.TransactionIndex++
	account.AllTransactions = append(account.AllTransactions, amount)
}

// Method to view transactions
func (account *BankAccount) ViewTransactions() {
	fmt.Println("Last 5 Transactions (Array):", account.Transactions)
	fmt.Println("All Transactions (Slice):", account.AllTransactions)
}

func main() {
	// Defer a goodbye message
	defer fmt.Println("Thank you for using the banking system. Goodbye!")

	// Map to store multiple accounts
	accounts := make(map[int]*BankAccount)

	// Create a sample account and add it to the map
	accounts[123456789] = &BankAccount{
		AccountHolder: "John Doe",
		AccountNumber: 123456789,
		Balance:       1000,
	}

	// Start a loop for multiple transactions
	for {
		// Display menu
		fmt.Println("\nBanking System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Transactions")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		// Get user choice
		var choice int
		fmt.Scan(&choice)

		// Handle user choice using a switch statement
		switch choice {
		case 1: // Deposit
			fmt.Print("Enter account number: ")
			var accountNumber int
			fmt.Scan(&accountNumber)

			account, exists := accounts[accountNumber]
			if !exists {
				fmt.Println("Account not found!")
				continue
			}

			fmt.Print("Enter amount to deposit: ")
			var deposit int
			fmt.Scan(&deposit)
			account.Deposit(deposit)

		case 2: // Withdraw
			fmt.Print("Enter account number: ")
			var accountNumber int
			fmt.Scan(&accountNumber)

			account, exists := accounts[accountNumber]
			if !exists {
				fmt.Println("Account not found!")
				continue
			}

			fmt.Print("Enter amount to withdraw: ")
			var withdraw int
			fmt.Scan(&withdraw)
			account.Withdraw(withdraw)

		case 3: // View Transactions
			fmt.Print("Enter account number: ")
			var accountNumber int
			fmt.Scan(&accountNumber)

			account, exists := accounts[accountNumber]
			if !exists {
				fmt.Println("Account not found!")
				continue
			}

			account.ViewTransactions()

		case 4: // Exit
			fmt.Println("Exiting the system...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

/*
type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "John", age: 30}
	p.greet()
	p.inc()

}

func (p *Person) inc() {
	p.age++
	fmt.Println("Age after increment:", p.age)
}
func (p *Person) greet() {
	fmt.Println("hello", p.name, p.age)

}

/*

func main() {
	value := 10

	fmt.Println("Value before change:", value)

	updatevalue(&value)
	fmt.Println("Value after change:", value)

}

func updatevalue(val *int) {
	*val = *val * 20
}

/*
func main() {
	salary := make(map[string]int)
	salary["John"] = 50000
	salary["Jane"] = 60000
	salary["Doe"] = 55000
	fmt.Println("Salary of :", salary["Jane"])

	dsalary, exists := salary["javed"]
	if exists {
		fmt.Println("Salary of Doe:", dsalary)
	} else {
		fmt.Println("Doe's salary not found")
	}
	salary["John"] = 70000 // Update John's salary
	fmt.Println("Updated salary of John:", salary["John"])

	fmt.Println("All salaries:")
	for name, dsalary := range salary {
		fmt.Printf("%s: %d\n", name, dsalary)
	}

	delete(salary, "John") // Delete John's salary
	delete(salary, "Jane") // Delete Jane's salary
	fmt.Println("After deletion, salaries:", salary)

}

/*
	// slices
	s := make([]int, 5)
	fmt.Println("Length of slice:", len(s), cap(s))

	s = append(s, 1)
	s = append(s, 2)
	fmt.Println("Slice after appending:", len(s), cap(s), s)
}

/*

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println((arr[i]))

	}
}

/*

type Author struct {
	fname string
	lname string
}

type Book struct {
	title  string
	page   int
	author Author
}

func main() {
	// Initialize an Author
	author := Author{fname: "John", lname: "Doe"}

	// Initialize a Book with the Author
	book := Book{title: "Go Programming", page: 300, author: author}

	// Print book details
	fmt.Println("Book Title:", book.title)
	fmt.Println("Number of Pages:", book.page)
	fmt.Println("Author Name:", book.author)
}

/*
func main() {
	// Initialize account balance using a pointer
	balance := 1000
	account := &balance

	// Defer a goodbye message
	defer fmt.Println("Thank you for using the banking system. Goodbye!")

	// Start a loop for multiple transactions
	for {
		// Display menu
		fmt.Println("\nBanking System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		// Get user choice
		var choice int
		fmt.Scan(&choice)

		// Handle user choice using a switch statement
		switch choice {
		case 1: // Deposit
			fmt.Print("Enter amount to deposit: ")
			var deposit int
			fmt.Scan(&deposit)
			if deposit > 0 {
				*account += deposit
				fmt.Println("Deposit successful! New balance:", *account)
			} else {
				fmt.Println("Invalid deposit amount.")
			}
		case 2: // Withdraw
			fmt.Print("Enter amount to withdraw: ")
			var withdraw int
			fmt.Scan(&withdraw)
			if withdraw > 0 && withdraw <= *account {
				*account -= withdraw
				fmt.Println("Withdrawal successful! New balance:", *account)
			} else if withdraw > *account {
				fmt.Println("Insufficient balance!")
			} else {
				fmt.Println("Invalid withdrawal amount.")
			}
		case 3: // Exit
			fmt.Println("Exiting the system...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

/*
	x := 10
	y := &x
	fmt.Println("Value of x:", x)          // Output: Value of x: 10
	fmt.Println("Address of x:", y)        // Output: Address of x: 0x...
	fmt.Println("Value at address y:", *y) // Output: Value at address y: 10
	*y = 20
	fmt.Println("New value of x:", x) // Output: New value of x: 20

}

/*
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Friday")
	}
}
*/
/*
func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println(("even"))
	} else {
		fmt.Println("odd")
	}
}
*/

/*
	func main() {
		fmt.Println("Hello, World!")
		a, b := 5, 3
		fmt.Println("Addition:", calc.Add(a, b))
		fmt.Println("Subtraction:", calc.Subtract(a, b))
		fmt.Println("Multiplication:", calc.Multiply(a, b))
		fmt.Println("Division:", calc.Divide(a, b))
		fmt.Println("Division by zero:", calc.Divide(a, 0))
	}
*/
