package main

import (
	"github.com/treyburn/boggle/repository"
	"github.com/treyburn/boggle/solver"
)

func buildCustomSolver(path string, repo repository.Repository) solver.Solver {
	var data [][]rune

	root := solver.NewTrie()
	for _, word := range data {
		root.Insert(word)
	}

	return &solver.OffTheShelf{}
}
