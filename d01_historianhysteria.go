// Solution to https://adventofcode.com/2024/day/1
package adventofcode

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func HistorianHysteria() Puzzle {
	return historianHysteria{}
}

type historianHysteria struct {
}

func (p historianHysteria) Details() Details {
	return Details{Day: 1, Description: "Historian Hysteria"}
}

func (p historianHysteria) Solve(input *Input) (Result, error) {
	left, right, err := p.parseInput(input)

	if err != nil {
		return Result{}, err
	}

	part1 := p.calcDistance(left, right)
	part2 := p.calcSimilarityScore(left, right)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}, nil
}

func (p historianHysteria) calcDistance(left, right []int) int {
	var distance int
	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return distance
}

func (p historianHysteria) calcSimilarityScore(left, right []int) int {
	var score int
	freq := make(map[int]int)

	for _, r := range right {
		freq[r] += 1
	}

	for _, l := range left {
		score += l * freq[l]
	}

	return score
}

func (p historianHysteria) parseInput(input *Input) ([]int, []int, error) {
	lineRgx := regexp.MustCompile(`^([0-9]+)   ([0-9]+)$`)
	var left, right []int

	for i, line := range input.Lines() {
		if !lineRgx.MatchString(line) {
			return nil, nil, fmt.Errorf("line %d is invalid: %s", i, line)
		}

		groups := lineRgx.FindAllStringSubmatch(line, -1)
		l, _ := strconv.Atoi(groups[0][1])
		r, _ := strconv.Atoi(groups[0][2])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right, nil
}
