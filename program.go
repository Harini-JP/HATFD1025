package main

import (
	"fmt"
)

// Function to find the maximum and second maximum values in an array
func findMaxAndSecondMax(arr []int) (int, int) {
	if len(arr) < 2 {
		fmt.Println("Array must have at least two elements")
		return -1, -1 // Edge case: if the array has less than 2 elements
	}

	// Initialize max and secondMax
	maxValue, secondMaxValue := arr[0], -1 

	
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxValue {
			secondMaxValue = maxValue 
			maxValue = arr[i]
		} else if arr[i] > secondMaxValue && arr[i] != maxValue {
			secondMaxValue = arr[i]
		}
	}

	// If no second maximum was found
	if secondMaxValue == -1 {
		fmt.Println("There is no second largest element")
		return maxValue, -1
	}

	return maxValue, secondMaxValue
}

func main() {
	// Sample input
	arr := []int{}
	maxValue, secondMaxValue := findMaxAndSecondMax(arr)

	if maxValue != -1 && secondMaxValue != -1 {
		fmt.Printf("The largest element is: %d\n", maxValue)
		fmt.Printf("The second largest element is: %d\n", secondMaxValue)
	}
}