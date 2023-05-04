package main

import "fmt"

func main() {
	var aNum, bNum, cNum float32

	fmt.Printf("Enter a:")
	fmt.Scanf("%f", &aNum)

	fmt.Printf("Enter b:")
	fmt.Scanf("%f", &bNum)

	fmt.Printf("Enter c:")
	fmt.Scanf("%f", &cNum)

	fmt.Print(aNum, bNum, cNum)
}
