package main

import (
	"fmt"
)

func RecurringNumber(array []int) int {

	var numCount []int = make([]int, 100)
	var numcountMap map[int]int = make(map[int]int)
	fmt.Println("Recurring number question")

	if array == nil {
		return 0
	}
	for _, val := range array {

		// if numCount[val] > 0  {
		if numcountMap[val] > 0 {
			return val
		} else {
			numCount[val] += 1
			numcountMap[val] += 1

		}

	}
	return 0
}
