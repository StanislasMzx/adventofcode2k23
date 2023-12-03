package main

import (
	_ "embed"
	"fmt"
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
	n1, p1, s1 := []int{}, []int{}, []int{}
	lines := strings.Split(f, "\n")
	lon := len(lines[0]) - 1
	for k, l := range lines {
		n2, p2, s2 := []int{}, []int{}, []int{}
		tmp := 0
		for i, c := range l {
			if unicode.IsDigit(c) {
				tmp = 10*tmp + int(c-'0')
			} else {
				if tmp != 0 {
					n2 = append(n2, tmp)
					p2 = append(p2, i-1)
					tmp = 0
				}
				if c != '.' {
					s2 = append(s2, i)
				}
			}
		}
		if tmp != 0 {
			n2 = append(n2, tmp)
			p2 = append(p2, lon)
		}
		r, to_rem := compareLines(n1, n2, p1, p2, s1, s2)
		res += r
		if k < len(lines)-1 {
			for _, i := range to_rem {
				n2[i] = 0
			}
			n1, p1, s1 = n2, p2, s2
		}
	}
	return res
}

func partTwo(f string) int {
	res := 0
	n1, p1, s1, m1 := []int{}, []int{}, []int{}, []int{}
	lines := strings.Split(f, "\n")
	lon := len(lines[0]) - 1
	for _, l := range lines {
		n2, p2, s2 := []int{}, []int{}, []int{}
		tmp := 0
		for i, c := range l {
			if unicode.IsDigit(c) {
				tmp = 10*tmp + int(c-'0')
			} else {
				if tmp != 0 {
					n2 = append(n2, tmp)
					p2 = append(p2, i-1)
					tmp = 0
				}
				if c == '*' {
					s2 = append(s2, i)
				}
			}
		}
		if tmp != 0 {
			n2 = append(n2, tmp)
			p2 = append(p2, lon)
		}
		m2 := make([]int, len(s2))
		for i := range m2 {
			m2[i] = 0
		}
		r, m2 := countAdjacents(n1, n2, p1, p2, s1, s2, m1, m2)
		res += r
		n1, p1, s1, m1 = n2, p2, s2, m2
	}
	return res
}

func compareLines(n1, n2, p1, p2, s1, s2 []int) (int, []int) {
	r := 0
	n := append(n1, n2...)
	p := append(p1, p2...)
	s := append(s1, s2...)
	to_rem := []int{}
	for i := range s {
		for j := range p {
			if s[i] >= p[j]-intLen(n[j]) && s[i] <= p[j]+1 {
				r += n[j]
				if j >= len(n1) {
					to_rem = append(to_rem, j-len(n1))
				}
			}
		}
	}
	return r, to_rem
}

func countAdjacents(n1, n2, p1, p2, s1, s2, m1, m2 []int) (int, []int) {
	r := 0
	n := append(n1, n2...)
	p := append(p1, p2...)
	s := append(s1, s2...)
	m := append(m1, m2...)
	for i := range s {
		for j := range p {
			if s[i] >= p[j]-intLen(n[j]) && s[i] <= p[j]+1 {
				if m[i] == 0 {
					m[i] = -n[j]
				} else if m[i] < 0 && -m[i] != n[j] {
					r += -m[i] * n[j]
					m[i] = 1
				}
			}
		}
	}
	return r, m[len(s1):]
}

func intLen(n int) int {
	l := 0
	for n != 0 {
		n /= 10
		l++
	}
	return l
}
