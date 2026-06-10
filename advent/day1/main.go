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

	for i, value := range values {
		fmt.Printf("Current Dial position: %d\n", start)
		fmt.Printf("Current count: %d\n", count)
		if value == "" {
			continue
		}
		fmt.Printf("Index: %d Value: %s\n", i, value)

		direction := string(value[0])
		fmt.Printf("d: %s\n",direction)

		if direction == "R" {
			n, err := convertInt(value)
			if err != nil {
				log.Fatal(err)
			}
			start += n
			fmt.Printf("combo: %d\n", start)
			calc := start % 100
			if calc ==  0{
				fmt.Printf("Calc: %d\n", calc)
				count++
			}
		}
		if direction == "L" {
			n, err := convertInt(value)
			fmt.Printf("rotateBy: %d\n", n)
			if err != nil {
				log.Fatal(err)
			}
			
			start -= n
			fmt.Printf("combo: %d\n", start)
			calc := start % 100
			if calc == 0 {
				fmt.Printf("Calc: %d\n", calc )
				count++
			}
		}
	}
	fmt.Printf("Password is: %d\n", count)
}
