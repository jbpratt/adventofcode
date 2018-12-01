package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day1.txt", "The input file to read")
var pt2 = flag.Bool("pt2", false, "Solve part two")

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	inputs := string(b)
	contents := strings.Split(inputs, "\n")
	result := 0
	dup := make(map[int]bool)
outer:
	for {
		for _, line := range contents {
			if len(line) == 0 {
				break
			}
			value, _ := strconv.Atoi(line)
			result += value
			if dup[result] && *pt2 {
				fmt.Printf("Repeated: %d\n", result)
				break outer
			}
			dup[result] = true
		}
		if !*pt2 {
			fmt.Printf("%d\n", result)
			break
		}
	}
}
