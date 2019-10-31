package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var size int

	fmt.Println("MERGE SORT\nInsert number of elements to sort: ")

	fmt.Scan(&size)

	array := generateArray(size)

	fmt.Println("\nVALUES:\n", array, "\n\nStarting... ")

	time.Sleep(time.Second * 3)

	fmt.Println("\nProcess:")

	fmt.Println("\nSORTED VALUES:\n", mergeSort(array))

}

func generateArray(size int) []int {

	array := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {

		array[i] = rand.Intn(size)

	}

	return array

}

func mergeSort(array []int) []int {

	if len(array) < 2 {

		return array

	}

	mid := (len(array)) / 2

	fmt.Println(" Divide:\t", array)

	return merge(mergeSort(array[:mid]), mergeSort(array[mid:]))

}

func merge(left, right []int) []int {

	size, resultIndex, leftIndex, rightIndex := len(left)+len(right), 0, 0, 0
	result := make([]int, size, size)

	for leftIndex < len(left) && rightIndex < len(right) {

		result, leftIndex, rightIndex = insert(result, left, right, resultIndex, leftIndex, rightIndex)

		resultIndex++

	}

	fill(result, left, resultIndex, leftIndex)
	fill(result, right, resultIndex, rightIndex)

	fmt.Println(" Conquer:\t", result)

	return result

}

func insert(result, left, right []int, resultIndex, leftIndex, rightIndex int) ([]int, int, int) {

	if left[leftIndex] <= right[rightIndex] {

		return insertion(result, left, resultIndex, leftIndex), leftIndex + 1, rightIndex

	}

	return insertion(result, right, resultIndex, rightIndex), leftIndex, rightIndex + 1

}

func insertion(main, secondary []int, mIndex, sIndex int) []int {

	main[mIndex] = secondary[sIndex]

	return main

}

func fill(result, rest []int, mainIndex, restIndex int) []int {

	if restIndex < len(rest) {

		result[mainIndex] = rest[restIndex]

		fill(result, rest, mainIndex+1, restIndex+1)

	}

	return result

}
