package dojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividerNumber1(t *testing.T) {
	assert := assert.New(t)
	numberToDivider := 8

	resp := dividerNumber(numberToDivider)
	expected := 4
	assert.Equal(resp, expected, "must be divider number 8 for 2 should return 4")
}

func TestDividerNumber2(t *testing.T) {
	assert := assert.New(t)
	numberToDivider := 10

	resp := dividerNumber(numberToDivider)
	expected := 5
	assert.Equal(resp, expected, "must be divider number 10 for 2 should return 5")
}
func TestDividerNumber3(t *testing.T) {
	assert := assert.New(t)
	numberToDivider := 2

	resp := dividerNumber(numberToDivider)
	expected := 1
	assert.Equal(resp, expected, "must be divider number 2 for 2 should return 1")
}

func TestDividerNumber4(t *testing.T) {
	assert := assert.New(t)
	numberToDivider := 100

	resp := dividerNumber(numberToDivider)
	expected := 50
	assert.Equal(resp, expected, "must be divider number 100 for 2 should return 50")
}

func TestRestOfTheDivision1(t *testing.T) {
	assert := assert.New(t)
	numberToForRestDivision := 100

	resp := restOfTheDivision(numberToForRestDivision)
	expected := 0
	assert.Equal(resp, expected, "must be rest of the division 100 for 2 should return 0")
}

func TestRestOfTheDivision2(t *testing.T) {
	assert := assert.New(t)
	numberToForRestDivision := 5

	resp := restOfTheDivision(numberToForRestDivision)
	expected := 1
	assert.Equal(resp, expected, "must be rest of the division 5 for 2 should return 1")
}
func TestRestOfTheDivision3(t *testing.T) {
	assert := assert.New(t)
	numberToForRestDivision := 25

	resp := restOfTheDivision(numberToForRestDivision)
	expected := 1
	assert.Equal(resp, expected, "must be rest of the division 25 for 2 should return 1")
}

func TestRestOfTheDivision4(t *testing.T) {
	assert := assert.New(t)
	numberToForRestDivision := 2

	resp := restOfTheDivision(numberToForRestDivision)
	expected := 0
	assert.Equal(resp, expected, "must be rest of the division 2 for 2 should return 0")
}

func TestReverseArray1(t *testing.T) {
	assert := assert.New(t)
	arrayToReverse := []int{1, 2}

	resp := reverseArray(arrayToReverse)
	expected := []int{2, 1}
	assert.Equal(resp, expected, "must be array to reverse [1,2] should return [2,1]")
}

func TestReverseArray2(t *testing.T) {
	assert := assert.New(t)
	arrayToReverse := []int{1, 2, 3}

	resp := reverseArray(arrayToReverse)
	expected := []int{3, 2, 1}
	assert.Equal(resp, expected, "must be array to reverse [1,2,3] should return [3,2,1]")
}

func TestReverseArray3(t *testing.T) {
	assert := assert.New(t)
	arrayToReverse := []int{1, 2, 3, 4}

	resp := reverseArray(arrayToReverse)
	expected := []int{4, 3, 2, 1}
	assert.Equal(resp, expected, "must be array to reverse [1,2,3,4] should return [4,3,2,1]")
}

func TestReverseArray4(t *testing.T) {
	assert := assert.New(t)
	arrayToReverse := []int{1, 2, 3, 4, 5, 6, 7}

	resp := reverseArray(arrayToReverse)
	expected := []int{7, 6, 5, 4, 3, 2, 1}
	assert.Equal(resp, expected, "must be array to reverse [1,2,3,4,5,6,7] should return [7,6,5,4,3,2,1]")
}

func TestBinary1(t *testing.T) {
	assert := assert.New(t)
	number := 2

	resp := binary(number)
	expected := []int{1, 0}
	assert.Equal(resp, expected, "must be number 2 should converte to value binary for [1,0]")
}

func TestBinary2(t *testing.T) {
	assert := assert.New(t)
	number := 8

	resp := binary(number)
	expected := []int{1, 0, 0}
	assert.Equal(resp, expected, "must be number 8 should converte to value binary for [1,0,0]")
}

func TestBinary3(t *testing.T) {
	assert := assert.New(t)
	number := 10

	resp := binary(number)
	expected := []int{1, 0, 1, 0}
	assert.Equal(resp, expected, "must be number 10 should converte to value binary for [1,0,1,0]")
}
