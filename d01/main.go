package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

var digits map[string]string = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	s := 0
	for _, l := range strings.Split(f, "\n") {
		i := extractDigits(l)
		if i == -1 {
			panic("Invalid input")
		}
		s += i
	}

	return s
}

func partTwo(f string) int {
	s := 0
	for _, l := range strings.Split(f, "\n") {
		s += extractDigits(remplaceStringDigits(l))
	}

	return s
}

func extractDigits(s string) int {
	f := func(c rune) bool {
		return unicode.IsDigit(c)
	}
	a := strings.IndexFunc(s, f)
	b := strings.LastIndexFunc(s, f)

	if a != -1 && b != -1 {
		return 10*int(s[a]-'0') + int(s[b]-'0')
	}
	return -1 // err
}

func remplaceStringDigits(s string) string {
	for k, v := range digits {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
