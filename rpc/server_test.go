package rpc

import (
	"context"
	"errors"
	"testing"

	"github.com/treyburn/boggle/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestBoggleService_Solve(t *testing.T) {
	ctx := context.Background()
	board := "a,b,c;a,b,c;a,b,c"
	req := &api.SolveRequest{Board: board}
	testLogger, err := zap.NewDevelopment()
	require.NoError(t, err)

	ms := &mockSolver{}
	ms.On("Solve", board)

	server := &BoggleService{
		solver: ms,
		repo:   nil,
		logger: testLogger,
	}

	resp, err := server.Solve(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetId())
}

func TestBoggleService_Solution(t *testing.T) {
	ctx := context.Background()
	id := "testid"
	wantSolution := []string{"test", "test", "test"}
	testLogger, err := zap.NewDevelopment()
	require.NoError(t, err)

	mr := &mockRepo{words: wantSolution}
	mr.On("Get", id).Return(nil, nil)

	server := &BoggleService{
		solver: nil,
		repo:   mr,
		logger: testLogger,
	}

	solution, err := server.Solution(ctx, &api.SolutionRequest{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, wantSolution, solution.Words)
}

func TestBoggleService_Solution_Err(t *testing.T) {
	ctx := context.Background()
	id := "testid"
	var wantSolution []string = nil
	testLogger, err := zap.NewDevelopment()
	require.NoError(t, err)

	mr := &mockRepo{words: wantSolution}
	mr.On("Get", id).Return(nil, errors.New("test error"))

	server := &BoggleService{
		solver: nil,
		repo:   mr,
		logger: testLogger,
	}

	solution, err := server.Solution(ctx, &api.SolutionRequest{Id: id})
	assert.Error(t, err)
	assert.Nil(t, solution)
}

type mockSolver struct {
	mock.Mock
}

func (m *mockSolver) Solve(_, board string) {
	_ = m.Called(board)
}

type mockRepo struct {
	mock.Mock
	words []string // doing this because testify/mock isn't quite type rich enough with args/return types
}

func (m *mockRepo) Get(id string) ([]string, error) {
	args := m.Called(id)
	return m.words, args.Error(1)
}

func (m *mockRepo) Put(id string, solution []string) {
	_ = m.Called(id, solution)
}

func (m *mockRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
