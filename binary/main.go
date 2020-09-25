package dojo

func dividerNumber(divider int) int {
	return divider / 2
}

func restOfTheDivision(divider int) int {
	return divider % 2
}

func reverseArray(arr []int) []int {
	lenght := len(arr) - 1

	for index := 0; index <= lenght/2; index++ {

		finalPositions := lenght - index
		aux := arr[finalPositions]
		arr[finalPositions] = arr[index]
		arr[index] = aux
	}

	return arr
}

func binary(numberToConvertBinary int) []int {
	numberBinary := []int{}

	for {
		dividend := dividerNumber(numberToConvertBinary)
		rest := restOfTheDivision(numberToConvertBinary)
		numberBinary = append(numberBinary, rest)
		numberToConvertBinary = dividend

		if dividend < 2 {
			numberBinary = append(numberBinary, dividend)
			break
		}
	}

	numberBinary = reverseArray(numberBinary)
	return numberBinary
}
