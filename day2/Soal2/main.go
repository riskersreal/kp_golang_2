package main

import "fmt"

func main() {
	var inputNum uint8
	var divisorNum uint8
	var sevenCheck bool = false

	fmt.Print("Enter a positive number:")
	fmt.Scanf("%d", &inputNum)

	divisorNum = 7

	if inputNum%divisorNum == 0 {
		sevenCheck = true
	}

	if sevenCheck == true {
		fmt.Println(inputNum, "is a number divisible by", divisorNum)
	} else {
		fmt.Println(inputNum, "is not a number divisible by", divisorNum)
	}
}
