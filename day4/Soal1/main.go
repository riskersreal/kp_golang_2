package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  float32
}

func (c *Car) calDistance() {
	calDistance := 1.5 * c.fuelIn
	fmt.Println("This car can travel for", calDistance, "mil")
}

func main() {
	var carInput Car

	fmt.Println("Masukkan jenis mobil")
	fmt.Scanf("%s\n", &carInput.typeCar)
	fmt.Println("Masukkan bahan bakar berapa liter (float)")
	fmt.Scanf("%f\n", &carInput.fuelIn)

	carInput.calDistance()
}
