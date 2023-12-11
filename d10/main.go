package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input []byte

var startX, startY, startPipe = 62, 111, byte('J')

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %d (in %s)\n", partOne(input), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d (in %s)\n", partTwo(input), time.Since(start))
}

func partOne(f []byte) int {
	steps, _ := getBoundaries(f)
	return steps
}

func partTwo(f []byte) int {
	_, area := getBoundaries(f)
	return area
}

func getBoundaries(f []byte) (int, int) {
	steps := 1
	area := 0
	lines := bytes.Split(f, []byte("\n"))
	linesR := len(lines)
	linesC := len(lines[0])
	x, y := 0, 0
	sketch := make([][]byte, linesR)
	for i := range sketch {
		sketch[i] = make([]byte, linesC)
	}
	boundaries := make([][]bool, linesR)
	for i := range boundaries {
		boundaries[i] = make([]bool, linesC)
	}

	for i, l := range lines {
		for j, c := range l {
			sketch[i][j] = c
			if c == 'S' {
				x, y = i, j
			}
		}
	}

	lastX, lastY := x, y
	end := false
	for !end {
		boundaries[x][y] = true
		tmpX, tmpY := x, y
		switch sketch[x][y] {
		case 'S':
			if x+1 < linesC && (sketch[x+1][y] == 'L' || sketch[x+1][y] == 'J' || sketch[x+1][y] == '|') {
				x++
				steps++
				continue
			} else if y+1 < linesR && (sketch[x][y+1] == '7' || sketch[x][y+1] == 'J' || sketch[x][y+1] == '-') {
				y++
				steps++
				continue
			} else if x-1 >= 0 && (sketch[x-1][y] == '7' || sketch[x-1][y] == 'F' || sketch[x-1][y] == '|') {
				x--
				steps++
				continue
			} else if y-1 >= 0 && (sketch[x][y-1] == 'L' || sketch[x][y-1] == 'F' || sketch[x][y-1] == '-') {
				y--
				steps++
				continue
			} else {
				panic("no next pipe")
			}
		case '|':
			if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == '|' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'L' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'J' {
				x++
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '|' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '7' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == 'F' {
				x--
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		case '-':
			if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '-' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == 'J' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '7' {
				y++
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == '-' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'L' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'F' {
				y--
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		case 'L':
			if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '|' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '7' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == 'F' {
				x--
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '-' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == 'J' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '7' {
				y++
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		case 'J':
			if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '|' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == '7' {
				x--
				steps++
			} else if x-1 >= 0 && lastX != x-1 && sketch[x-1][y] == 'F' {
				x--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == '-' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'L' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'F' {
				y--
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		case '7':
			if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == '|' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'L' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'J' {
				x++
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == '-' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'L' {
				y--
				steps++
			} else if y-1 >= 0 && lastY != y-1 && sketch[x][y-1] == 'F' {
				y--
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		case 'F':
			if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == '|' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'L' {
				x++
				steps++
			} else if x+1 < linesR && lastX != x+1 && sketch[x+1][y] == 'J' {
				x++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '-' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == 'J' {
				y++
				steps++
			} else if y+1 < linesC && lastY != y+1 && sketch[x][y+1] == '7' {
				y++
				steps++
			} else {
				end = true
			}
			lastX, lastY = tmpX, tmpY
		}
	}

	// replace S with the fitting pipe
	sketch[startX][startY] = startPipe

	inside := false
	lastChar := byte('.')
	for i := range boundaries {
		inside = false
		lastChar = byte('.')
		for j := range boundaries[i] {
			if boundaries[i][j] {
				if sketch[i][j] == '-' {
					continue
				} else if sketch[i][j] == '|' || sketch[i][j] == 'L' || sketch[i][j] == 'F' {
					inside = !inside
				} else {
					switch lastChar {
					case '|':
					case '.':
						inside = !inside
					case 'L':
						if sketch[i][j] == 'J' {
							inside = !inside
						}
					case 'F':
						if sketch[i][j] == '7' {
							inside = !inside
						}
					}
				}
				lastChar = sketch[i][j]
			} else {
				if inside && !boundaries[i][j] {
					lastChar = '.'
					area++
				}
			}
		}
	}

	return steps / 2, area
}
