package adventofcode

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolutions(t *testing.T) {
	tests := []struct {
		puzzle Puzzle
		want   Result
	}{
		{
			puzzle: HistorianHysteria{},
			want: Result{
				Part1: "TODO",
				Part2: "TODO",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.puzzle.Details().String(), func(t *testing.T) {
			input, err := LoadInput(tc.puzzle)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got, err := tc.puzzle.Solve(&input)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("unexpected diff (-want +got):\n%s", diff)
			}
		})
	}
}
