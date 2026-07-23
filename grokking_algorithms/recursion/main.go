package main

import (
	"fmt"
	"time"
)

func countdown(i int) {
	fmt.Printf("%d\n", i)

	if i <= 1 {
		return
	} else {
		time.Sleep(1 * time.Second)
		countdown(i - 1)
	}
}

func main() {
	countdown(3)
}
