package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func convertInt(v string) (int, error) {
	i, err := strconv.Atoi(v[1:])
	if err != nil {
		log.Fatal(err)
	}
	return i, nil
}

func main(){
	count := 0
	start := 50
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	values := strings.Split(string(data), "\n")

	for _, value := range values {
		if value == "" {
			continue
		}

		direction := string(value[0])
		rotateBy, err := convertInt(value)
		if err != nil {
			log.Fatal(err)
		}

		for i:=0; i<rotateBy; i++ {
			if direction == "L" {
				start -= 1
			} else {
				start += 1
			}
			if start % 100 == 0 {
				count++
			}
		}
	}
	fmt.Printf("Password is: %d\n", count)
}
