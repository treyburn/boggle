package rpc

import (
	"context"

	"github.com/treyburn/boggle/api"
	"github.com/treyburn/boggle/repository"
	"github.com/treyburn/boggle/solver"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BoggleService struct {
	api.UnimplementedBoggleServiceServer
	solver solver.Solver
	repo   repository.Repository
	logger *zap.Logger
}

func (bs *BoggleService) Solve(ctx context.Context, req *api.SolveRequest) (*api.SolveResponse, error) {
	bs.logger.Info("received solve request", zap.String("board", req.GetBoard()))
	id := uuid.New().String()
	// maybe we should set the status here - so we could have a status api
	// we could extend the repo interface - or perhaps make it generic and have multiple repos
	go bs.solver.Solve(id, req.GetBoard())
	bs.logger.Info("created solve process", zap.String("id", id))
	return &api.SolveResponse{Id: id}, nil
}

func (bs *BoggleService) Solution(_ context.Context, req *api.SolutionRequest) (*api.SolutionResponse, error) {
	bs.logger.Info("received solution request", zap.String("id", req.GetId()))
	// if we had the status api - we could check that first and error if status wasn't completed
	solution, err := bs.repo.Get(req.GetId())
	if err != nil {
		bs.logger.Error("failed to get solution", zap.Error(err))
		return nil, err
	}
	return &api.SolutionResponse{Words: solution}, nil
}

func NewBoggleService(repo repository.Repository, solver solver.Solver, logger *zap.Logger) *BoggleService {
	return &BoggleService{
		solver: solver,
		repo:   repo,
		logger: logger,
	}
}
