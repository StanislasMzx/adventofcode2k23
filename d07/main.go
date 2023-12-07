package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type hand struct {
	cards string
	bid   int
}

func main() {
	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}

func partOne(f string) int {
	res := 0
	var hands []hand
	strength := map[uint8]int{
		'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,
	}
	lines := strings.Split(f, "\n")

	for _, l := range lines {
		fields := strings.Fields(l)
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{fields[0], bid})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands, strength, false, i, j)
	})

	for k, v := range hands {
		res += (k + 1) * v.bid
	}
	return res
}

func partTwo(f string) int {
	res := 0
	var hands []hand
	strength := map[uint8]int{
		'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
	}
	lines := strings.Split(f, "\n")

	for _, l := range lines {
		fields := strings.Fields(l)
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{fields[0], bid})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands, strength, true, i, j)
	})

	for k, v := range hands {
		res += (k + 1) * v.bid
	}
	return res
}

func getType(s string, joker bool) int {
	occ := [5]int{0, 0, 0, 0, 0}
	idx := map[uint8]int{}
	matched := 0
	nbj := 0

	for i := 0; i < 5; i++ {
		if joker && s[i] == 'J' {
			nbj++
			continue
		}
		for j := 0; j <= i; j++ {
			if s[i] == s[j] {
				k, ok := idx[s[i]]
				if !ok {
					idx[s[i]] = matched
					matched++
				}
				k = idx[s[i]]
				occ[k]++
				break
			}
		}
	}

	if joker {
		m := 0
		for i := 0; i < 5; i++ {
			if occ[i] > occ[m] {
				m = i
			}
		}
		occ[m] += nbj
	}

	if occ == [5]int{1, 1, 1, 1, 1} {
		return 1
	} else if occ == [5]int{2, 1, 1, 1, 0} || occ == [5]int{1, 2, 1, 1, 0} || occ == [5]int{1, 1, 2, 1, 0} || occ == [5]int{1, 1, 1, 2, 0} {
		return 2
	} else if occ == [5]int{2, 2, 1, 0, 0} || occ == [5]int{2, 1, 2, 0, 0} || occ == [5]int{1, 2, 2, 0, 0} {
		return 3
	} else if occ == [5]int{3, 1, 1, 0, 0} || occ == [5]int{1, 3, 1, 0, 0} || occ == [5]int{1, 1, 3, 0, 0} {
		return 4
	} else if occ == [5]int{3, 2, 0, 0, 0} || occ == [5]int{2, 3, 0, 0, 0} {
		return 5
	} else if occ == [5]int{4, 1, 0, 0, 0} || occ == [5]int{1, 4, 0, 0, 0} {
		return 6
	} else {
		return 7
	}
}

func compareHands(h []hand, s map[uint8]int, joker bool, i, j int) bool {
	a := getType(h[i].cards, joker)
	b := getType(h[j].cards, joker)
	if a < b {
		return true
	} else if a > b {
		return false
	}

	for k := 0; k < 5; k++ {
		if s[h[i].cards[k]] == s[h[j].cards[k]] {
			continue
		}
		if s[h[i].cards[k]] < s[h[j].cards[k]] {
			return true
		} else {
			return false
		}
	}
	return false
}
