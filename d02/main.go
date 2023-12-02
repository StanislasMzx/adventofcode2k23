package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type config struct {
	r int
	g int
	b int
}

func main() {
	c := config{r: 12, g: 13, b: 14}
	fmt.Printf("Part one: %d\n", partOne(input, &c))
	fmt.Printf("Part two: %d\n", partTwo(input))
}

func partOne(f string, c *config) int {
	s := 0
	for i, l := range strings.Split(f, "\n") {
		_, l, _ := strings.Cut(l, ":")
		configs := strToConfigs(l)
		if compareConfigs(configs, *c) {
			s += i + 1
		}
	}
	return s
}

func partTwo(f string) int {
	s := 0
	for _, l := range strings.Split(f, "\n") {
		_, l, _ := strings.Cut(l, ":")
		configs := strToConfigs(l)
		s += findMinSetPower(configs)
	}
	return s
}

func strToConfigs(s string) *[]config {
	var configs []config
	for _, c := range strings.Split(s, ";") {
		var config config
		for _, v := range strings.Split(c, ",") {
			num, col := 0, ""
			fmt.Sscanf(v, "%d %s", &num, &col)
			switch col {
			case "red":
				config.r = num
			case "green":
				config.g = num
			case "blue":
				config.b = num
			}
		}
		configs = append(configs, config)
	}
	return &configs
}

func compareConfigs(c *[]config, d config) bool {
	for _, config := range *c {
		if config.r > d.r || config.g > d.g || config.b > d.b {
			return false
		}
	}
	return true
}

func findMinSetPower(c *[]config) int {
	r := config{r: 0, g: 0, b: 0}
	for _, conf := range *c {
		if conf.r > r.r {
			r.r = conf.r
		}
		if conf.g > r.g {
			r.g = conf.g
		}
		if conf.b > r.b {
			r.b = conf.b
		}
	}
	return r.r * r.g * r.b
}
