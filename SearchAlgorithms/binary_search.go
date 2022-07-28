package main

import (
	"fmt"
)

func BinarySearch() {
	array := []int{3, 4, 6, 8, 9, 34, 78, 100}
	var x int
	fmt.Println("Given array ", array, "\n enter number to search ")
	fmt.Scanf("%d", &x)
	fmt.Println("searching ", x)
	search(array, x)
	searchIndex(array, 0, len(array)-1, x)
}

/*
Search if a number exists in a given array
*/
func search(arr []int, x int) {

	if len(arr) == 0 {
		fmt.Println("Not Found")
		return
	}
	start := 0
	end := len(arr) - 1
	mid := (start + end) / 2
	if x < arr[mid] {
		search(arr[:mid], x)
	} else if x > arr[mid] {
		search(arr[mid+1:], x)
	} else {
		fmt.Println("Found ", x)
		return
	}

}

/*
Search index of the searching element
*/
func searchIndex(arr []int, s int, e int, x int) {

	if e >= s {

		mid := s + (e-s)/2
		fmt.Println("end : ", e, " start : ", s, " mid : ", mid)

		if arr[mid] == x {
			fmt.Println("Found at index ", mid)
			return
		} else if arr[mid] > x {
			searchIndex(arr, s, mid-1, x)
		} else {
			searchIndex(arr, mid+1, e, x)
		}

	}
}
