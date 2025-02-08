package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	// fmt.Scanln(&name)

	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello, ", name)
	fmt.Println("Sum of 5 and 6 is: ", add(5, 6))
}

func add(a int, b int) int {
	return a + b
}
