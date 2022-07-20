package solver

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildBoard(t *testing.T) {
	type testCase struct {
		name          string
		strBoard      string
		wantBoardXLen int
		wantBoardYLen int
		wantErr       error
	}

	var tests = []testCase{
		{"valid", "a,b,c;a,b,c;a,b,c", 3, 3, nil},
		{"rectangle", "a,b,c;a,b,c", 0, 0, errors.New("board is not square")},
		{"rectangle", "a,b,c;a,,c;a,b,c", 0, 0, errors.New("invalid character in row: a,,c")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			board, err := buildRuneBoard(tc.strBoard)
			assert.Equal(t, err, tc.wantErr)
			if tc.wantErr == nil {
				assert.Equal(t, tc.wantBoardYLen, len(board))
				assert.Equal(t, tc.wantBoardXLen, len(board[0]))
			}
		})
	}
}

func Test_getEdges(t *testing.T) {
	type testCase struct {
		name      string
		board     [][]rune
		x         int
		y         int
		wantEdges []rune
	}

	var tests = []testCase{
		{"3x3 - middle", [][]rune{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, 1, 1, []rune{1, 2, 3, 1, 3, 1, 2, 3}},
		{"3x3 - upper center", [][]rune{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, 1, 0, []rune{1, 3, 1, 2, 3}},
		{"3x3 - upper left", [][]rune{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, 0, 0, []rune{2, 1, 2}},
		{"3x3 - lower right", [][]rune{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, 2, 2, []rune{2, 3, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := getEdges(tc.x, tc.y, tc.board)

			assert.Equal(t, tc.wantEdges, got)
		})
	}
}
