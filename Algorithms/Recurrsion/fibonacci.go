package recurrsion

func GetFibonacciSeries(number int) int {

	if number < 2 {
		return number
	}
	return GetFibonacciSeries(number-1) + GetFibonacciSeries(number-2)
}
