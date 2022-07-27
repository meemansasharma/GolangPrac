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
	duplicateList := recurrsion.Node{}
	head := duplicateList.CreateList()
	fmt.Println(duplicateList.CheckDuplicates(head))
	duplicateList.Insert(head, 2)
	duplicateList.Insert(head, 2)
	duplicateList.Insert(head, 2)
	duplicateList.Insert(head, 4)
	duplicateList.Insert(head, 6)
	duplicateList.Print(head)
	duplicateList.CheckDuplicates(head)
	duplicateList.Print(head)
}
