package main

func MergeSort(Array1 []int, Array2 []int) []int {

	if Array1 == nil {
		return Array2
	}
	if Array2 == nil {
		return Array1
	}

	var mergedArray []int
	i, j := 0, 0
	for i < len(Array1) && j < len(Array2) {
		if Array1[i] < Array2[j] {
			mergedArray = append(mergedArray, Array1[i])
			i++
		} else {
			mergedArray = append(mergedArray, Array2[j])
			j++
		}

	}
	if i != len(Array1) {
		mergedArray = append(mergedArray, Array1[i:]...)
	} else if j != len(Array2) {
		mergedArray = append(mergedArray, Array2[j:]...)
	}
	return mergedArray
}
