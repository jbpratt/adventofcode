package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day1.txt", "The input file to read")

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	inputs := string(b)
	contents := strings.Split(inputs, "\n")
	result := 0
	for _, line := range contents {
		if len(line) == 0 {
			break
		}
		value, _ := strconv.Atoi(line)
		result += value
	}
	fmt.Printf("%d\n", result)
}
