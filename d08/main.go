package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type lr struct {
	left  string
	right string
}

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	lines := strings.Split(f, "\n")
	nav := map[string]lr{}

	for _, l := range lines[2:] {
		nav[l[0:3]] = lr{l[7:10], l[12:15]}
	}

	return pathLen(lines, nav, "AAA", "ZZZ")
}

func partTwo(f string) int {
	lines := strings.Split(f, "\n")
	nav := map[string]lr{}
	s := []string{}
	paths := []int{}

	for _, l := range lines[2:] {
		k := l[0:3]
		nav[k] = lr{l[7:10], l[12:15]}
		if k[2] == 'A' {
			s = append(s, k)
		}
	}

	for _, v := range s {
		paths = append(paths, pathLen(lines, nav, v, "Z"))
	}

	return lcm(paths[0], paths[1], paths[2:]...)
}

func pathLen(lines []string, nav map[string]lr, s, e string) int {
	res := 0
	dir := lines[0]

	i := 0
	l := len(dir)

	if len(e) == 1 {
		for string(s[2]) != e {
			if dir[i] == 'L' {
				s = nav[s].left
			} else {
				s = nav[s].right
			}
			i = (i + 1) % l
			res++
		}
	} else {
		for s != e {
			if dir[i] == 'L' {
				s = nav[s].left
			} else {
				s = nav[s].right
			}
			i = (i + 1) % l
			res++
		}
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, tab ...int) int {
	res := a * b / gcd(a, b)

	for i := 0; i < len(tab); i++ {
		res = lcm(res, tab[i])
	}

	return res
}
