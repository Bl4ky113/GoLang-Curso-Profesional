package main

import "fmt"

func main() {
	var number1 int
	var number2 int
	fmt.Print("Enter Numbers: ")
	fmt.Scanln(&number1, &number2)
	fmt.Printf("Numbers: %d, %d\n", number1, number2)
	fmt.Printf("Addition of the Numbers: %d\n", number1+number2)
	fmt.Printf("1st Subtraction of the numbers: %d\n", number1-number2)
	fmt.Printf("2nd Subtraction of the numbers: %d\n", number2-number1)
	fmt.Printf("Product of the numbers: %d\n", number1*number2)
	fmt.Printf("1st Division of the numbers: %d\nresidue: %d\n", number1/number2, number1%number2)
	fmt.Printf("2nd Division of the numbers: %d\nresidue: %d\n", number2/number1, number1%number2)

	var value_product float32
	var taxPercentage float32 = 18.0
	var costumer_value_product float32

	fmt.Print("Enter a Product Value: ")
	fmt.Scanln(&value_product)
	costumer_value_product = value_product * (1.0 + (taxPercentage / 100.0))
	fmt.Printf("Costumer Product Value: %f", costumer_value_product)
}
