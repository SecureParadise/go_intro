package main

// Our main aim is to understand data types in golang
import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var name string = "John Doe"
	var age int = 25
	var isMarried bool = false
	var height float32 = 5.8
	var weight float32 = 70.5
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Married: ", isMarried)
	fmt.Println("Height: ", height)
	fmt.Println("Weight: ", weight)
}
