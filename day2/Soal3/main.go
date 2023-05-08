package main

import "fmt"

func main() {
	var baseOneNum float32
	var baseTwoNum float32
	var heightNum float32
	var areaNum float32

	fmt.Printf("Enter the size of base 1:")
	fmt.Scanf("%f\n", &baseOneNum)

	fmt.Printf("Enter the size of base 2:")
	fmt.Scanf("%f\n", &baseTwoNum)

	fmt.Printf("Enter the size of height:")
	fmt.Scanf("%f\n", &heightNum)

	areaNum = 0.5 * (baseOneNum + baseTwoNum) * heightNum

	fmt.Print("The area of this trapezium is:", areaNum)
}
