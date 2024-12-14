// Solution to https://adventofcode.com/2024/day/3
package adventofcode

import (
	"regexp"
	"strconv"
)

func MullItOver() Puzzle {
	return mullItOver{}
}

type mullItOver struct {
}

func (p mullItOver) Details() Details {
	return Details{Day: 3, Description: "Mull It Over"}
}

func (p mullItOver) Solve(input *Input) (Result, error) {
	mulls := p.parseInput(input)
	part1 := p.addUp(mulls)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: "",
	}, nil
}

func (p mullItOver) addUp(mulls []mull) int {
	var res int

	for _, mull := range mulls {
		res += mull.result()
	}

	return res
}

func (p mullItOver) parseInput(input *Input) []mull {
	rgx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var res []mull

	for _, line := range input.Lines() {
		for _, group := range rgx.FindAllStringSubmatch(line, -1) {
			n1, _ := strconv.Atoi(group[1])
			n2, _ := strconv.Atoi(group[2])
			res = append(res, mull{n1, n2})
		}
	}

	return res
}

type mull struct {
	n1, n2 int
}

func (m mull) result() int {
	return m.n1 * m.n2
}
