package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func is_even(value string) bool {
	if len(value)%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	sum := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for idRange := range strings.SplitSeq(string(data), ",") {
		idRange = strings.TrimSpace(idRange)
		if idRange == "" {
			continue
		}
		split := strings.Split(idRange, "-")
		first_bound, err1 := strconv.Atoi(split[0])
		last_bound, err2 := strconv.Atoi(split[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error1: %s\n Error2: %s", err1, err2)
			log.Fatal("Failed to parse bounds")
		}
		for i := first_bound; i <= last_bound; i++ {
			idString := strconv.Itoa(i)

			pattern := regexp2.MustCompile(`^(\d+)\1+$`, 0)

			m, err := pattern.FindStringMatch(idString)
			if err != nil {
				log.Fatal(err)
			}

			if m != nil {
				sum += i
			}
			// PART 1
			// if m {
			// 	sum += i
			// }
			// if len(idString)%2 == 0 {
			// 	half := len(idString) / 2
			// 	if idString[:half] == idString[half:] {
			// 		sum += i
			// 	}
			// }
		}
	}
	fmt.Printf("Sum: %d", sum)
}
