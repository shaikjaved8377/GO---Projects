package main

import "fmt"

func main() {
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
