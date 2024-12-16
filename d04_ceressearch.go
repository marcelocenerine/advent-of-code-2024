// Solution to https://adventofcode.com/2024/day/4
package adventofcode

import (
	"fmt"
	"strconv"
)

func CeresSearch() Puzzle {
	return ceresSearch{}
}

type ceresSearch struct {
}

func (p ceresSearch) Details() Details {
	return Details{Day: 4, Description: "Ceres Search"}
}

func (p ceresSearch) Solve(input *Input) (Result, error) {
	matrix, err := p.parseInput(input)

	if err != nil {
		return Result{}, err
	}

	part1 := p.countOccurrences([]rune("XMAS"), matrix)
	return Result{
		Part1: strconv.Itoa(part1),
		Part2: "TODO",
	}, nil
}

func (p ceresSearch) countOccurrences(word []rune, matrix [][]rune) int {
	var count int

	for r, row := range matrix {
		for c := 0; c < len(row); c++ {
			for horizDir := -1; horizDir <= 1; horizDir++ {
				for vertDir := -1; vertDir <= 1; vertDir++ {
					count += p.search(word, 0, matrix, r, c, horizDir, vertDir)
				}
			}
		}
	}

	return count
}

func (p ceresSearch) search(word []rune, wordIdx int, matrix [][]rune, row, col, horizDir, vertDir int) int {
	if (horizDir == 0 && vertDir == 0) ||
		row < 0 || row >= len(matrix) ||
		col < 0 || col >= len(matrix[0]) ||
		wordIdx >= len(word) ||
		word[wordIdx] != matrix[row][col] {
		return 0
	}

	if wordIdx == len(word)-1 {
		return 1
	}

	return p.search(word, wordIdx+1, matrix, row+horizDir, col+vertDir, horizDir, vertDir)
}

func (p ceresSearch) parseInput(input *Input) ([][]rune, error) {
	var res [][]rune

	for i, line := range input.Lines() {
		if length := len(line); i > 0 && length != len(res[0]) {
			return nil, fmt.Errorf("line %d has invalid length: %d", i, length)
		}

		res = append(res, []rune(line))
	}

	return res, nil
}
