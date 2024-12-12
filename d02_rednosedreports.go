// Solution to https://adventofcode.com/2024/day/2
package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

type RedNosedReports struct{}

func (p RedNosedReports) Details() Details {
	return Details{Day: 2, Description: "Red-Nosed Reports"}
}

func (p RedNosedReports) Solve(input *Input) (Result, error) {
	reports, err := p.parseInput(input)

	if err != nil {
		return Result{}, err
	}

	part1 := p.countSafe(reports)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: "TODO",
	}, nil
}

func (p RedNosedReports) countSafe(reports []Report) int {
	var count int

outer:
	for _, rep := range reports {
		if len(rep.levels) > 1 {
			increasing := rep.levels[0] < rep.levels[1]

			for i := 1; i < len(rep.levels); i++ {
				curr := rep.levels[i]
				prev := rep.levels[i-1]
				diff := math.Abs(float64(prev - curr))

				if diff < 1 || diff > 3 ||
					(prev < curr && !increasing) ||
					(prev > curr && increasing) {
					continue outer
				}
			}
		}

		count += 1
	}

	return count
}

func (p RedNosedReports) parseInput(input *Input) ([]Report, error) {
	var reports []Report

	for i, line := range input.Lines() {
		parts := strings.Split(line, " ")
		var levels []int

		for _, part := range parts {
			level, err := strconv.Atoi(part)

			if err != nil {
				return nil, fmt.Errorf("line %d is invalid: %s. %w", i, line, err)
			}

			levels = append(levels, level)
		}

		reports = append(reports, Report{levels})
	}

	return reports, nil
}
