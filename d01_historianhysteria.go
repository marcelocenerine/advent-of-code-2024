// Solution to https://adventofcode.com/2024/day/1
package adventofcode

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

var lineRgx = regexp.MustCompile(`^([0-9]+)   ([0-9]+)$`)

type HistorianHysteria struct{}

func (p HistorianHysteria) Details() Details {
	return Details{Day: 1, Description: "Historian Hysteria"}
}

func (p HistorianHysteria) Solve(input *Input) (Result, error) {
	l1, l2, err := p.parseInput(input)

	if err != nil {
		return Result{}, err
	}

	part1 := p.calculateDistance(l1, l2)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: "TODO",
	}, nil
}

func (p HistorianHysteria) calculateDistance(l1, l2 []int) int {
	sort.Ints(l1)
	sort.Ints(l2)
	distance := 0

	for i := 0; i < len(l1); i++ {
		distance += int(math.Abs(float64(l1[i] - l2[i])))
	}

	return distance
}

func (p HistorianHysteria) parseInput(input *Input) ([]int, []int, error) {
	var list1 []int
	var list2 []int

	for i, line := range input.Lines() {
		if !lineRgx.MatchString(line) {
			return nil, nil, fmt.Errorf("line %d is invalid: %s", i, line)
		}

		groups := lineRgx.FindAllStringSubmatch(line, -1)
		col1, _ := strconv.Atoi(groups[0][1])
		col2, _ := strconv.Atoi(groups[0][2])
		list1 = append(list1, col1)
		list2 = append(list2, col2)
	}

	return list1, list2, nil
}
