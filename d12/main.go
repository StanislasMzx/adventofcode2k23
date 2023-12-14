package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed input.txt
var input []byte

type arrang struct {
	springs      string
	groups       string
	currentGroup int
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %d (in %s)\n", partOne(input), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d (in %s)\n", partTwo(input), time.Since(start))
}

func partOne(f []byte) int {
	res := 0
	mem = make(map[arrang]int)

	for _, l := range bytes.Split(f, []byte("\n")) {
		fields := bytes.FieldsFunc(l, func(r rune) bool { return r == ',' || r == ' ' })
		groups := make([]int, len(fields)-1)
		for i, g := range fields[1:] {
			groups[i], _ = strconv.Atoi(string(g))
		}
		res += arrangements(fields[0], groups, 0)
	}

	return res
}

func partTwo(f []byte) int {
	res := 0
	mem = make(map[arrang]int)

	for _, l := range bytes.Split(f, []byte("\n")) {
		fields := bytes.FieldsFunc(l, func(r rune) bool { return r == ',' || r == ' ' })
		groups := make([]int, len(fields)-1)
		for i, g := range fields[1:] {
			groups[i], _ = strconv.Atoi(string(g))
		}

		springs5 := make([]byte, len(fields[0])*5+4)
		for i := 0; i < 5; i++ {
			copy(springs5[i*(len(fields[0])+1):], fields[0])
			if i < 4 {
				springs5[(i+1)*(len(fields[0])+1)-1] = '?'
			}
		}
		groups5 := make([]int, len(groups)*5)
		for i := 0; i < 5; i++ {
			copy(groups5[i*len(groups):], groups)
		}

		res += arrangements(springs5, groups5, 0)
	}

	return res
}

func arrangements(springs []byte, groups []int, currentGroup int) int {
	if i, ok := mem[toArrang(springs, groups, currentGroup)]; ok {
		return i
	} else {
		if len(groups) == 1 && groups[0] == currentGroup && !bytes.Contains(springs, []byte("#")) {
			return setMem(springs, groups, currentGroup, 1)
		} else if len(groups) == 0 || len(springs) == 0 {
			return setMem(springs, groups, currentGroup, 0)
		} else if currentGroup == groups[0] {
			switch springs[0] {
			case '#':
				return setMem(springs, groups, currentGroup, 0)
			default:
				res := arrangements(springs[1:], groups[1:], 0)
				return setMem(springs, groups, currentGroup, res)
			}
		} else if currentGroup < groups[0] {
			switch springs[0] {
			case '#':
				res := arrangements(springs[1:], groups, currentGroup+1)
				return setMem(springs, groups, currentGroup, res)
			case '.':
				if currentGroup == 0 {
					res := arrangements(springs[1:], groups, 0)
					return setMem(springs, groups, currentGroup, res)
				} else {
					return setMem(springs, groups, currentGroup, 0)
				}
			case '?':
				if currentGroup == 0 {
					res := arrangements(springs[1:], groups, 0) + arrangements(springs[1:], groups, 1)
					return setMem(springs, groups, currentGroup, res)
				} else {
					res := arrangements(springs[1:], groups, currentGroup+1)
					return setMem(springs, groups, currentGroup, res)
				}
			}
		}
	}
	return setMem(springs, groups, currentGroup, 0)
}

var mem = make(map[arrang]int)

func setMem(springs []byte, groups []int, currentGroup, nbArrang int) int {
	mem[toArrang(springs, groups, currentGroup)] = nbArrang
	return nbArrang
}

func toArrang(springs []byte, groups []int, currentGroup int) arrang {
	return arrang{string(springs), fmt.Sprintf("%v", groups), currentGroup}
}
