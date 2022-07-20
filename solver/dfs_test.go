package solver

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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

func TestDFS_search(t *testing.T) {
	wantSolutionLength := 6
	board := "a,b,c;d,a,a;d,t,t"

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	data := [][]rune{[]rune("bat"), []rune("cat"), []rune("tac"), []rune("tab"), []rune("abc")}

	root := NewTrie()
	for _, word := range data {
		root.Insert(word)
	}

	rb, err := buildRuneBoard(board)
	require.NoError(t, err)
	graph := buildGraph(rb)

	repo := newRepoSpy()
	dfs := NewDfs(root, repo, logger)

	got, err := dfs.search(graph)
	assert.NoError(t, err)
	assert.Equal(t, wantSolutionLength, len(got))
}

func TestDFS_Solve(t *testing.T) {
	testId := "test"
	wantSolutionLength := 6
	board := "a,b,c;d,a,a;d,t,t"

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	data := [][]rune{[]rune("bat"), []rune("cat"), []rune("tac"), []rune("tab"), []rune("abc")}

	root := NewTrie()
	for _, word := range data {
		root.Insert(word)
	}

	repo := newRepoSpy()
	dfs := NewDfs(root, repo, logger)

	dfs.Solve(testId, board)
	
	got, err := repo.Get(testId)
	assert.NoError(t, err)
	assert.Equal(t, wantSolutionLength, len(got))
}