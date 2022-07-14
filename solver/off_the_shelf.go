package solver

import (
	"strings"

	"github.com/gammazero/bogglesolver/solver"
	"go.uber.org/zap"
)

type Repository interface {
	Put(id string, solutions []string)
}

type OffTheShelf struct {
	repo       Repository
	dictionary string
	logger     *zap.Logger
}

func (ots *OffTheShelf) Solve(id, board string) {
	b, err := parseBoard(board)
	if err != nil {
		ots.logger.Error("issue parsing board", zap.Error(err))
		return
	}
	boggleSolver, err := solver.NewSolver(b.xLen, b.yLen, ots.dictionary, true)
	if err != nil {
		ots.logger.Error("issue creating solver", zap.Error(err))
		return
	}
	solution, err := boggleSolver.Solve(b.board)
	if err != nil {
		ots.logger.Error("issue while solving", zap.Error(err))
		return
	}

	ots.repo.Put(id, solution)
}

func NewOffTheShelf(repo Repository, logger *zap.Logger, filepath string) *OffTheShelf {
	return &OffTheShelf{
		repo:       repo,
		dictionary: filepath,
		logger:     logger,
	}
}

type boardInfo struct {
	xLen  int
	yLen  int
	board string
}

func parseBoard(board string) (boardInfo, error) {
	y := strings.Count(board, ";") + 1
	x := (strings.Count(board, ",") / y) + 1
	b := strings.Replace(board, ";", "", -1)
	b = strings.Replace(b, ",", "", -1)
	return boardInfo{xLen: x, yLen: y, board: b}, nil
}
