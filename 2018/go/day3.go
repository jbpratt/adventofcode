package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day3.txt", "File to use")

func Split(r rune) bool {
	return r == '@' || r == ':' || r == ',' || r == 'x'
}

func main() {
	flag.Parse()
	file, err := os.Open(*inputFile)
	if err != nil {
		return
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var input []string
	for scan.Scan() {
		input = append(input, scan.Text())
	}
	if scan.Err() != nil {
		fmt.Println("Finished..")
	}

	var grid [1000][1000]int
	var total int
	for i := 0; i < len(input)-1; i++ {
		current := input[i]
		piece := BreakUp(current)
		xPos, _ := strconv.Atoi(strings.TrimSpace(piece[1]))
		yPos, _ := strconv.Atoi(strings.TrimSpace(piece[2]))
		xDim, _ := strconv.Atoi(strings.TrimSpace(piece[3]))
		yDim, _ := strconv.Atoi(strings.TrimSpace(piece[4]))
		for j := 0; j < xDim; j++ {
			for k := 0; k < yDim; k++ {
				if grid[yPos+k][xPos+j] == 1 {
					total++
				}
				grid[yPos+k][xPos+j]++
			}
		}
	}
	fmt.Println(total)
	for i := 0; i < len(input)-1; i++ {
		var ol = false
		current := input[i]
		piece := BreakUp(current)
		xPos, _ := strconv.Atoi(strings.TrimSpace(piece[1]))
		yPos, _ := strconv.Atoi(strings.TrimSpace(piece[2]))
		xDim, _ := strconv.Atoi(strings.TrimSpace(piece[3]))
		yDim, _ := strconv.Atoi(strings.TrimSpace(piece[4]))
		for j := 0; j < xDim; j++ {
			for k := 0; k < yDim; k++ {
				if grid[yPos+k][xPos+j] > 1 {
					ol = true
				}
			}
		}
		if ol == false {
			fmt.Println(piece[0])
		}
	}
}

func BreakUp(a string) []string {
	c := strings.FieldsFunc(a, Split)
	return c
}

// --- Day 3: No Matter How You Slice It ---

// The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit (thanks to someone who helpfully wrote its box IDs on the wall of the warehouse in the middle of the night). Unfortunately, anomalies are still affecting them - nobody can even agree on how to cut the fabric.

// The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side.

// Each Elf has made a claim about which area of fabric would be ideal for Santa's suit. All claims have an ID and consist of a single rectangle with edges parallel to the edges of the fabric. Each claim's rectangle is defined as follows:

//     The number of inches between the left edge of the fabric and the left edge of the rectangle.
//     The number of inches between the top edge of the fabric and the top edge of the rectangle.
//     The width of the rectangle in inches.
//     The height of the rectangle in inches.

// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall. Visually, it claims the square inches of fabric represented by # (and ignores the square inches of fabric represented by .) in the diagram below:

// ...........
// ...........
// ...#####...
// ...#####...
// ...#####...
// ...#####...
// ...........
// ...........
// ...........

// The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas. For example, consider the following claims:

// #1 @ 1,3: 4x4
// #2 @ 3,1: 4x4
// #3 @ 5,5: 2x2

// Visually, these claim the following areas:

// ........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........

// The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while adjacent to the others, does not overlap either of them.)

// If the Elves all proceed with their own plans, none of them will have enough fabric. How many square inches of fabric are within two or more claims?
