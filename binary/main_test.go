package dojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividerNumber1(t *testing.T) {
	assert := assert.New(t)
	divider := 8

	resp := dividerNumber(divider)
	expected := 4
	assert.Equal(resp, expected, "must be divider number 8 for 2 should return 4")
}

func TestDividerNumber2(t *testing.T) {
	assert := assert.New(t)
	divider := 10

	resp := dividerNumber(divider)
	expected := 5
	assert.Equal(resp, expected, "must be divider number 10 for 2 should return 5")
}
func TestDividerNumber3(t *testing.T) {
	assert := assert.New(t)
	divider := 2

	resp := dividerNumber(divider)
	expected := 1
	assert.Equal(resp, expected, "must be divider number 2 for 2 should return 1")
}

func TestDividerNumber4(t *testing.T) {
	assert := assert.New(t)
	divider := 100

	resp := dividerNumber(divider)
	expected := 50
	assert.Equal(resp, expected, "must be divider number 100 for 2 should return 50")
}

func TestRestOfTheDivision1(t *testing.T) {
	assert := assert.New(t)
	divider := 100

	resp := restOfTheDivision(divider)
	expected := 0
	assert.Equal(resp, expected, "must be rest of the division 100 for 2 should return 0")
}

func TestRestOfTheDivision2(t *testing.T) {
	assert := assert.New(t)
	divider := 5

	resp := restOfTheDivision(divider)
	expected := 1
	assert.Equal(resp, expected, "must be rest of the division 5 for 2 should return 1")
}
func TestRestOfTheDivision3(t *testing.T) {
	assert := assert.New(t)
	divider := 25

	resp := restOfTheDivision(divider)
	expected := 1
	assert.Equal(resp, expected, "must be rest of the division 25 for 2 should return 1")
}

func TestRestOfTheDivision4(t *testing.T) {
	assert := assert.New(t)
	divider := 2

	resp := restOfTheDivision(divider)
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
	numberToConvertBinary := 2

	resp := binary(numberToConvertBinary)
	expected := []int{1, 0}
	assert.Equal(resp, expected, "must be number 2 should converte to value binary for [1,0]")
}

func TestBinary2(t *testing.T) {
	assert := assert.New(t)
	numberToConvertBinary := 4

	resp := binary(numberToConvertBinary)
	expected := []int{1, 0, 0}
	assert.Equal(resp, expected, "must be number 4 should converte to value binary for [1,0,0]")
}

func TestBinary3(t *testing.T) {
	assert := assert.New(t)
	numberToConvertBinary := 10

	resp := binary(numberToConvertBinary)
	expected := []int{1, 0, 1, 0}
	assert.Equal(resp, expected, "must be number 10 should converte to value binary for [1,0,1,0]")
}

func TestBinary4(t *testing.T) {
	assert := assert.New(t)
	numberToConvertBinary := 15

	resp := binary(numberToConvertBinary)
	expected := []int{1, 1, 1, 1}
	assert.Equal(resp, expected, "must be number 15 should converte to value binary for [1,1,1,1]")
}
