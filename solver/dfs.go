package solver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/treyburn/boggle/repository"

	"go.uber.org/zap"
)

type DFS struct {
	dictionary *Trie
	repo       repository.Repository
	logger     *zap.Logger
}

func NewDfs(dict *Trie, repo repository.Repository, log *zap.Logger) *DFS {
	return &DFS{
		dictionary: dict,
		repo:       repo,
		logger:     log,
	}
}

func (d *DFS) Solve(id string, board string) {
	rb, err := buildRuneBoard(board)
	if err != nil {
		d.logger.Error("issue building board", zap.Error(err), zap.String("ID", id))
		return
	}
	graph := buildGraph(rb)

	answers, err := d.search(graph)
	if err != nil {
		d.logger.Error("issue searching for solution", zap.Error(err), zap.String("ID", id))
	}

	d.repo.Put(id, answers)
}

func (d *DFS) search(graph [][]*Node) ([]string, error) {
	answers := make([]string, 0)
	for _, row := range graph {
		for _, node := range row {
			word := []rune{node.char}
			if !d.dictionary.IsPrefix(word) {
				continue
			}
			if d.dictionary.IsWord(word) {
				answers = append(answers, string(word))
			}
			edges := getEdges(node, graph)
			for len(edges) > 0 {
				word = append(word, edges[0].char)
				if !d.dictionary.IsPrefix(word) {
					// backtrack word
					word = word[:len(word)-1]
					edges = edges[1:]
					continue
				}
				if d.dictionary.IsWord(word) {
					answers = append(answers, string(word))
				}
				subEdges := getEdges(edges[0], graph)
				subEdges = append(subEdges, edges[1:]...)
				edges = subEdges
			}
		}
	}
	return answers, nil
}

func buildRuneBoard(board string) ([][]rune, error) {
	y := strings.Count(board, ";") + 1
	x := (strings.Count(board, ",") / y) + 1
	if x != y {
		return nil, errors.New("board is not square")
	}
	var b = make([][]rune, 0)
	rows := strings.Split(board, ";")
	for _, row := range rows {
		trimmed := strings.ReplaceAll(row, ",", "")
		tiles := []rune(trimmed)
		if len(tiles) != x {
			return nil, errors.New(fmt.Sprintf("invalid character in row: %v", row))
		}
		b = append(b, tiles)
	}
	return b, nil
}
