package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type seedMap struct {
	dr  int
	sr  int
	len int
}

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	s := []int{}

	lines := strings.Split(f, "\n")
	fields := strings.Fields(lines[0])
	for _, e := range fields[1:] {
		i, _ := strconv.Atoi(e)
		s = append(s, i)
	}
	lines = append(lines, "EOF")

	return parseMaps(lines, s)
}

func partTwo(f string) int {
	s := []int{}

	lines := strings.Split(f, "\n")
	fields := strings.Fields(lines[0])
	incr := -1
	for _, e := range fields[1:] {
		if incr == -1 {
			i, _ := strconv.Atoi(e)
			s = append(s, i)
			incr = i
		} else {
			i, _ := strconv.Atoi(e)
			for j := 1; j < i; j++ {
				s = append(s, incr+j)
			}
			incr = -1
		}
	}
	lines = append(lines, "EOF")

	return parseMaps(lines, s)
}

func parseMaps(lines []string, s []int) int {
	maps := []seedMap{}

	for _, l := range lines[2:] {
		// fmt.Println(s)
		if l == "" {
			continue
		}
		if unicode.IsDigit(rune(l[0])) {
			fields := strings.Fields(l)
			dr, _ := strconv.Atoi(fields[0])
			sr, _ := strconv.Atoi(fields[1])
			len, _ := strconv.Atoi(fields[2])
			maps = append(maps, seedMap{dr, sr, len})
		} else if len(maps) > 0 {
			s = mapSeeds(s, maps)
			maps = []seedMap{}
		}
	}

	res := s[0]
	for _, e := range s[1:] {
		if e < res {
			res = e
		}
	}

	return res
}

func mapSeeds(s []int, m []seedMap) []int {
	res := make([]int, len(s))
	copy(res, s)
	for n, i := range s {
		for _, e := range m {
			if i >= e.sr && i <= e.sr+e.len-1 {
				res[n] = e.dr + (i - e.sr)
				break
			}
		}
	}
	return res
}
