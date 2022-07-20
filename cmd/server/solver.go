package main

import (
	"github.com/treyburn/boggle/reader"
	"github.com/treyburn/boggle/repository"
	"github.com/treyburn/boggle/solver"
	"go.uber.org/zap"
)

func buildCustomSolver(path string, repo repository.Repository, logger *zap.Logger) (solver.Solver, error) {
	data, err := reader.Read(path)
	if err != nil {
		return nil, err
	}

	root := solver.NewTrie()
	for _, word := range data {
		rWord := []rune(word)
		root.Insert(rWord)
	}

	return solver.NewDfs(root, repo, logger), nil
}
