package main

import "fmt"

func main() {

	fmt.Println("Enter a number")
	var num int
	fmt.Scan(&num)
	result := 1

	for i := 1; i <= num; i++ {
		result = result * i

	}

	fmt.Println("The number is :", result)

}
