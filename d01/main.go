package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

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
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(dat)
	file.Split(bufio.ScanLines)

	s := 0
	for file.Scan() {
		i := extractDigits(file.Text())
		if i == -1 {
			panic("Invalid input")
		}
		s += i
	}

	// fmt.Printf("Sum: %d\n", s)
	dat.Close()
}

func extractDigits(s string) int {
	s = remplaceStringDigits(s)

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
