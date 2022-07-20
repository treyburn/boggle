package solver

import (
	"strings"

	"github.com/treyburn/boggle/repository"

	"github.com/gammazero/bogglesolver/solver"
	"go.uber.org/zap"
)

type OffTheShelf struct {
	repo       repository.Repository
	dictionary string
	logger     *zap.Logger
}

func (ots *OffTheShelf) Solve(id, board string) {
	b, err := parseBoard(board)
	if err != nil {
		ots.logger.Error("issue parsing board", zap.Error(err), zap.String("ID", id))
		return
	}
	boggleSolver, err := solver.NewSolver(b.xLen, b.yLen, ots.dictionary, true)
	if err != nil {
		ots.logger.Error("issue creating solver", zap.Error(err), zap.String("ID", id))
		return
	}
	solution, err := boggleSolver.Solve(b.board)
	if err != nil {
		ots.logger.Error("issue while solving", zap.Error(err), zap.String("ID", id))
		return
	}

	ots.repo.Put(id, solution)
}

func NewOffTheShelf(filepath string, repo repository.Repository, logger *zap.Logger) *OffTheShelf {
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
