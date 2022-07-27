package main

import (
	recurrsion "MimmiTraining/practice/Algorithms/Recurrsion"
	"fmt"
)

func main() {

	fmt.Println("Factorial of a number -", recurrsion.GetFactorial(5))
	//0,1,1,2,3,5
	fmt.Println("Factorial of a number -", recurrsion.GetFibonacciSeries(4))
	fmt.Println(recurrsion.ReverseString("meemansa"))

}
