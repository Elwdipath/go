package main

import (
	"fmt"
	"os"
	"strconv"
)

func binary_search(arr []int, to_search int) bool {
	found := false
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == to_search {
			found = true
			return found
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error converting to int")
	}
	result := binary_search([]int{1, 4, 8, 16, 32, 64}, guess)
	fmt.Println(result)
}
