package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day2.txt", "File to use")
var pt2 = flag.Bool("pt2", false, "To run part two of the problem")

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	answer := solve(string(file))
	fmt.Println(answer)
}

func solve(s string) string {
	Two := 0
	Three := 0

	file := strings.Split(s, "\n")
	for _, id := range file {
		Count := make(map[rune]int)
		foundTwo := false
		foundThree := false
		for _, l := range id {
			Count[l]++
		}
		for _, c := range Count {
			if c == 2 && !foundTwo {
				Two++
				foundTwo = true
			} else if c == 3 && !foundThree {
				Three++
				foundThree = true
			}
			if foundTwo && foundThree {
				break
			}
		}

	}

	return strconv.Itoa(Two * Three)
}
