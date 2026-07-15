package main
import (
	"fmt"
)


func findSmallest(arr []int) int{
	smallest := arr[0]
	smallest_index := 0
	
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func main () {
	start_index := findSmallest([]int{4,2,9,22})
	fmt.Printf("Smallest index is: %d\n", start_index)
}
