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
	part1 := p.addUpMulls(input)
	part2 := p.addUpEnabledMulls(input)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}, nil
}

func (p mullItOver) addUpMulls(input *Input) int {
	rgx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var res int

	for _, line := range input.Lines() {
		for _, group := range rgx.FindAllStringSubmatch(line, -1) {
			n1, _ := strconv.Atoi(group[1])
			n2, _ := strconv.Atoi(group[2])
			res += n1 * n2
		}
	}

	return res
}

func (p mullItOver) addUpEnabledMulls(input *Input) int {
	rgx := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	res := 0
	enabled := true

	for _, line := range input.Lines() {
		for _, group := range rgx.FindAllStringSubmatch(line, -1) {
			if group[0] == "do()" {
				enabled = true
				continue
			}

			if group[0] == "don't()" {
				enabled = false
				continue
			}

			if !enabled {
				continue
			}

			n1, _ := strconv.Atoi(group[1])
			n2, _ := strconv.Atoi(group[2])
			res += n1 * n2
		}
	}

	return res
}
