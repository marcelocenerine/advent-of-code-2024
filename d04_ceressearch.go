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

	part1 := p.countXmasOcurrences(matrix)
	part2 := p.countXShapedMasOccurrences(matrix)
	return Result{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}, nil
}

func (p ceresSearch) countXmasOcurrences(matrix [][]rune) int {
	var count int
	word := []rune("XMAS")

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			for horizDir := -1; horizDir <= 1; horizDir++ {
				for vertDir := -1; vertDir <= 1; vertDir++ {
					if p.search(word, 0, matrix, r, c, horizDir, vertDir) {
						count += 1
					}
				}
			}
		}
	}

	return count
}

func (p ceresSearch) countXShapedMasOccurrences(matrix [][]rune) int {
	var count int
	word := []rune("MAS")

	for r := 1; r < len(matrix)-1; r++ {
		for c := 1; c < len(matrix[r])-1; c++ {
			diag1 :=
				p.search(word, 0, matrix, r-1, c-1, 1, 1) || // top-down, left-right
					p.search(word, 0, matrix, r+1, c+1, -1, -1) // bottom-up, right-left

			if !diag1 {
				continue
			}

			diag2 :=
				p.search(word, 0, matrix, r-1, c+1, 1, -1) || // top-down, right-left
					p.search(word, 0, matrix, r+1, c-1, -1, 1) // bottom-up, left-right

			if diag2 {
				count += 1
			}
		}
	}

	return count
}

func (p ceresSearch) search(word []rune, wordIdx int, matrix [][]rune, row, col, horizDir, vertDir int) bool {
	if (horizDir == 0 && vertDir == 0) ||
		row < 0 || row >= len(matrix) ||
		col < 0 || col >= len(matrix[0]) ||
		wordIdx >= len(word) ||
		word[wordIdx] != matrix[row][col] {
		return false
	}

	if wordIdx == len(word)-1 {
		return true
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
