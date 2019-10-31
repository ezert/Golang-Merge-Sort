package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	size := 9

	array := generateValues(size)

	fmt.Println("Unsorted:\t", array)

	fmt.Println("Sorted:  \t", mergeSort(array))

}

func generateValues(size int) []int {

	array := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(array); i++ {

		array[i] = rand.Intn(size)

	}

	return array

}

func mergeSort(array []int) []int {

	if len(array) < 2 {

		return array

	}

	mid := len(array) / 2

	return merge(mergeSort(array[mid:]), mergeSort(array[:mid]))

}

func merge(left, right []int) []int {

	size, arrayIndex, leftIndex, rightIndex := len(left)+len(right), 0, 0, 0
	array := make([]int, size, size)

	for leftIndex < len(left) && rightIndex < len(right) {

		array, leftIndex, rightIndex = insertion(array, left, right, arrayIndex, leftIndex, rightIndex)

		arrayIndex++

	}

	if leftIndex < len(left) {

		array = fill(array, left, arrayIndex, leftIndex)

		return array

	}

	if rightIndex < len(right) {

		array = fill(array, right, arrayIndex, rightIndex)

	}

	return array

}

func insertion(array, left, right []int, arrayIndex, leftIndex, rightIndex int) ([]int, int, int) {

	if left[leftIndex] < right[rightIndex] {

		return insert(array, left, arrayIndex, leftIndex), leftIndex + 1, rightIndex

	}

	return insert(array, right, arrayIndex, rightIndex), leftIndex, rightIndex + 1

}

func insert(to, from []int, i, j int) []int {

	to[i] = from[j]

	return to

}

func fill(to, from []int, i, j int) []int {

	if j < len(from) {

		to = insert(to, from, i, j)

		fill(to, from, i+1, j+1)

	}

	return to

}
