// Solution to https://adventofcode.com/2024/day/2
package adventofcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RedNosedReports() Puzzle {
	return redNosedReports{}
}

type redNosedReports struct{}

func (p redNosedReports) Details() Details {
	return Details{Day: 2, Description: "Red-Nosed Reports"}
}

type report []int

func (p redNosedReports) Solve(input *Input) (Result, error) {
	reports, err := p.parseInput(input)

	if err != nil {
		return Result{}, err
	}

	part1 := p.countSafe(reports)
	part2 := p.countSafeWithTolerance(reports)

	return Result{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}, nil
}

func (p redNosedReports) countSafe(reports []report) int {
	var count int

	for _, rp := range reports {
		if p.isSafe(rp) {
			count += 1
		}
	}

	return count
}

func (p redNosedReports) countSafeWithTolerance(reports []report) int {
	var count int

	for _, rp := range reports {
		for i := 0; i < len(rp); i++ {
			var tmp report
			tmp = append(tmp, rp[:i]...)
			tmp = append(tmp, rp[i+1:]...)

			if p.isSafe(tmp) {
				count += 1
				break
			}
		}
	}

	return count
}

func (p redNosedReports) isSafe(rep report) bool {
	if len(rep) <= 1 {
		return true
	}

	increasing := rep[0] < rep[1]

	for i := 1; i < len(rep); i++ {
		curr := rep[i]
		prev := rep[i-1]
		diff := math.Abs(float64(prev - curr))

		if diff < 1 || diff > 3 || (prev < curr && !increasing) || (prev > curr && increasing) {
			return false
		}
	}

	return true
}

func (p redNosedReports) parseInput(input *Input) ([]report, error) {
	var reports []report

	for i, line := range input.Lines() {
		parts := strings.Split(line, " ")
		var report []int

		for _, part := range parts {
			level, err := strconv.Atoi(part)

			if err != nil {
				return nil, fmt.Errorf("line %d is invalid: %s. %w", i, line, err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
