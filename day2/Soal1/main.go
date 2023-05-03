package main

import "fmt"

func main() {
	var inputNum uint8
	var primeCheck bool = true
	var i uint8 = 2

	fmt.Print("Enter a positive number:")
	fmt.Scanf("%d", &inputNum)

	for i < inputNum && primeCheck == true {
		if inputNum%i == 0 {
			primeCheck = false
		}
		i++
	}

	if primeCheck == true {
		fmt.Println(inputNum, "is a prime number")
	} else {
		fmt.Println(inputNum, "is not a prime number")
	}
}
