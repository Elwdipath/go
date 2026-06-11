package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"strconv"
)

func is_even(value int) bool {
	if value % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	productIds := strings.Split(string(data), ",")

	for _, idRange := range productIds {
		idRange = strings.TrimSpace(idRange)
		if idRange == "" {
			continue
		}
		split := strings.Split(idRange, "-")
		first_bound,err1 := strconv.Atoi(split[0])
		last_bound, err2 := strconv.Atoi(split[1])
		
		if err1 != nil || err2 != nil {
			fmt.Printf("Error1: %s\n Error2: %s", err1, err2)
			log.Fatal("Failed to parse bounds")
		}
		for i:=first_bound; i<=last_bound; i++ {
			
			id := is_even(i)

		}
	}
}
