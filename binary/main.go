package dojo

func dividerNumber(number int) int {
	return number / 2
}

func restOfTheDivision(number int) int {
	return number % 2
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

func binary(number int) []int {
	resp := []int{}

	for {
		dividend := dividerNumber(number)
		rest := restOfTheDivision(number)
		resp = append(resp, rest)

		if dividend < 2 {
			resp = append(resp, dividend)
			break
		}
	}

	resp = reverseArray(resp)
	return resp
	// if number == 10 {
	// 	return []int{1, 0, 1, 0}
	// }
	// if number == 8 {
	// 	return []int{1, 0, 0}
	// }
	// return []int{1, 0}
}
