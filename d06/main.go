package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	res := 1
	l := strings.Split(f, "\n")
	t := strings.Fields(l[0][5:])
	d := strings.Fields(l[1][9:])
	k := len(t)
	for i := 0; i < k; i++ {
		b, _ := strconv.Atoi(t[i])
		c, _ := strconv.Atoi(d[i])
		x1, x2 := recordMargin(-1, b, -c)
		res *= x2 - x1 + 1
	}
	return res
}

func partTwo(f string) int {
	res := 1
	l := strings.Split(f, "\n")
	t := strings.ReplaceAll(l[0][5:], " ", "")
	d := strings.ReplaceAll(l[1][9:], " ", "")
	b, _ := strconv.Atoi(t)
	c, _ := strconv.Atoi(d)
	x1, x2 := recordMargin(-1, b, -c)
	res *= x2 - x1 + 1
	return res
}

func recordMargin(a, b, c int) (int, int) {
	d := b*b - 4*a*c
	if d == 0 {
		x := -b / (2 * a)
		return x, x
	} else if d > 0 {
		dr := math.Sqrt(float64(d))
		x1 := (-float64(b) - dr) / float64(2*a)
		x2 := (-float64(b) + dr) / float64(2*a)
		if x1 < x2 {
			if math.Trunc(x1) == x1 {
				x1++
			}
			if math.Trunc(x2) == x2 {
				x2++
			}
			return int(math.Ceil(x1)), int(math.Floor(x2))
		} else {
			if math.Trunc(x1) == x1 {
				x1--
			}
			if math.Trunc(x2) == x2 {
				x2++
			}
			return int(math.Ceil(x2)), int(math.Floor(x1))
		}
	}
	return -1, -1
}
