package recurrsion

func factorialOfANumber(number int) int {

	if number == 1 {
		return number
	}
	return number * factorialOfANumber((number - 1))
}

func GetFactorial(number int) int {
	return factorialOfANumber(number)
}
