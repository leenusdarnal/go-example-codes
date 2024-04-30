package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var size int
	println("Enter size of array")
	fmt.Scan(&size)

	var arr = make([]int, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
	}

	for i := 0; i < len(arr); i++ {
	}
}
