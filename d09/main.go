package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %d (in %s)\n", partOne(input), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d (in %s)\n", partTwo(input), time.Since(start))
}

func partOne(f string) int {
	res := 0
	tmp := 0

	for _, l := range strings.Split(f, "\n") {
		fields := strings.Fields(l)
		var diff []int
		a, _ := strconv.Atoi(fields[0])
		n := 0
		for i := 1; i < len(fields); i++ {
			n, _ = strconv.Atoi(fields[i])
			diff = append(diff, n-a)
			a = n
		}
		tmp += n
		tmp += diff[len(diff)-1]

		for !sliceIsNull(diff) {
			for i := 1; i < len(diff); i++ {
				diff[i-1] = diff[i] - diff[i-1]
			}
			diff = diff[:len(diff)-1]
			tmp += diff[len(diff)-1]
		}

		res += tmp
		tmp = 0
	}

	return res
}

func partTwo(f string) int {
	res := 0
	tmp := 0

	for _, l := range strings.Split(f, "\n") {
		fields := strings.Fields(l)
		var diff []int
		a, _ := strconv.Atoi(fields[0])
		tmp += a
		n := 0
		for i := 1; i < len(fields); i++ {
			n, _ = strconv.Atoi(fields[i])
			diff = append(diff, n-a)
			a = n
		}
		tmp -= diff[0]

		p := 1
		for !sliceIsNull(diff) {
			for i := 1; i < len(diff); i++ {
				diff[i-1] = diff[i] - diff[i-1]
			}
			diff = diff[:len(diff)-1]
			tmp += p * diff[0]
			p *= -1
		}

		res += tmp
		tmp = 0
	}

	return res
}

func sliceIsNull(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
