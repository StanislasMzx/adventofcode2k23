package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	res := 0
	for _, l := range strings.Split(f, "\n") {
		for i, c := range l {
			if c == ':' {
				w, j := parseWinning(l[i+2:])
				res += countPoints(l[i+j+4:], w)
			}
		}
	}
	return res
}

func partTwo(f string) int {
	res := 0
	lines := strings.Split(f, "\n")
	t := make([]int, len(lines))
	for i := range t {
		t[i] = 1
	}
	for n, l := range lines {
		for i, c := range l {
			if c == ':' {
				w, j := parseWinning(l[i+2:])
				res = countPoints(l[i+j+4:], w)
				if res > 0 {
					res = int(math.Log2(float64(res))) + 1
					for k := 1; k < res+1; k++ {
						t[n+k] += t[n]
					}
				}
			}
		}
	}
	for _, v := range t {
		res += v
	}
	return res
}

func parseWinning(s string) ([]int, int) {
	r := []int{}
	tmp := 0
	for i, c := range s {
		if unicode.IsDigit(c) {
			tmp = 10*tmp + int(c-'0')
		} else if c == ' ' {
			if tmp != 0 {
				r = append(r, tmp)
				tmp = 0
			}
		} else {
			if tmp != 0 {
				r = append(r, tmp)
				tmp = 0
			}
			return r, i
		}
	}
	return r, 0
}

func countPoints(s string, w []int) int {
	p := 0
	tmp := 0
	lon := len(s)
	for i, c := range s {
		if unicode.IsDigit(c) {
			tmp = 10*tmp + int(c-'0')
			if i == lon-1 {
				if tmp != 0 {
					if slices.Contains(w, tmp) {
						if p == 0 {
							p = 1
						} else {
							p *= 2
						}
					}
					tmp = 0
				}
			}
		} else if c == ' ' {
			if tmp != 0 {
				if slices.Contains(w, tmp) {
					if p == 0 {
						p = 1
					} else {
						p *= 2
					}
				}
				tmp = 0
			}
		} else {
			return p
		}
	}
	return p
}
