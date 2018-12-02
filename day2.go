package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day2.txt", "File to use")

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}
	strArray := strings.Fields(string(file))
	answer1 := solve(strArray)
	answer2 := solve2(strArray)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func common(a string, b string) string {
	var c string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			c += string(a[i])
		}
	}
	return c
}

// part 2
func solve2(strArray []string) string {
	var id string
	for i, a := range strArray {
		for j, b := range strArray {
			if i != j {
				temp := common(a, b)
				if len(temp) > len(id) {
					id = temp
				}
			}
		}
	}
	return id
}

// part 1
func solve(file []string) string {
	Two := 0
	Three := 0
	for _, id := range file {
		Count := make(map[rune]int)
		foundTwo, foundThree := false, false
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
