package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input []byte

type coord struct {
	x, y int
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %d (in %s)\n", partOne(input), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d (in %s)\n", partTwo(input), time.Since(start))
}

func partOne(f []byte) int {
	res := 0

	galaxies := expandSpace(f, 1)

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			res += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}

	return res
}

func partTwo(f []byte) int {
	res := 0

	galaxies := expandSpace(f, 999999)

	// compute the shortest distance (manhattan) between each pair of galaxies. Do not count the galaxies twice
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			res += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}

	return res
}

func expandSpace(f []byte, e int) []coord {
	lines := bytes.Split(f, []byte("\n"))
	linesX, linesY := len(lines), len(lines[0])
	var galaxies []coord
	EmptyRows := make([]bool, linesX)
	nonEmptyCols := make([]bool, linesY)

	addedRows := 0
	for i, l := range lines {
		empty := true
		for j, c := range l {
			if c == '#' {
				galaxies = append(galaxies, coord{i + addedRows*e, j})
				nonEmptyCols[j] = true
				empty = false
			}
		}
		EmptyRows[i] = empty
		if empty {
			addedRows++
		}
	}

	addedCols := 0
	for j, c := range nonEmptyCols {
		if !c {
			for i := range galaxies {
				if galaxies[i].y >= j+addedCols*e {
					galaxies[i].y += e
				}
			}
			addedCols++
		}
	}

	return galaxies
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
