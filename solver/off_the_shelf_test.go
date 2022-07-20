package solver

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestOffTheShelfSolver_Solve(t *testing.T) {
	wantSolutionLength := 14
	board := "a,b,c;d,a,a;d,t,t"

	testDictionary, err := filepath.Abs("../assets/3_letter_dictionary.txt")
	require.NoError(t, err)
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	repo := newRepoSpy()
	ots := NewOffTheShelf(testDictionary, repo, logger)

	ots.Solve("testid", board)
	solution, _ := repo.Get("testid")
	assert.Equal(t, wantSolutionLength, len(solution))
}

type repoSpy struct {
	cache map[string][]string
}

func (rs *repoSpy) Put(id string, solution []string) {
	rs.cache[id] = solution
}

func (rs *repoSpy) Get(id string) ([]string, error) {
	return rs.cache[id], nil
}

func (rs *repoSpy) Delete(id string) error {
	return nil
}

func newRepoSpy() *repoSpy {
	c := make(map[string][]string, 0)

	return &repoSpy{cache: c}
}
