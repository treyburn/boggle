package solver

import (
	"errors"
	"fmt"
	"github.com/treyburn/boggle/repository"
	"go.uber.org/zap"
	"strings"
)

type DFS struct {
	dictionary *Trie
	repo       repository.Repository
	logger     zap.Logger
}

func NewDfs(dict *Trie, repo repository.Repository, log zap.Logger) *DFS {
	return &DFS{
		dictionary: dict,
		repo:       repo,
		logger:     log,
	}
}

func (d *DFS) Solve(id string, board string) {
	b, err := buildRuneBoard(board)
	if err != nil {
		d.logger.Error("issue building board", zap.Error(err), zap.String("ID", id))
		return
	}

	answers, err := d.search(b)
	if err != nil {
		d.logger.Error("issue searching for solution", zap.Error(err), zap.String("ID", id))
	}

	d.repo.Put(id, answers)
}

func (d *DFS) search(board [][]rune) ([]string, error) {
	answers := make([]string, 0)
	for yIdx, row := range board {
		for xIdx, char := range row {
			word := []rune{char}
			if !d.dictionary.IsPrefix(word) {
				continue
			}
			if d.dictionary.IsWord(word) {
				answers = append(answers, string(word))
			}
			edges := getEdges(xIdx, yIdx, board)
			for len(edges) > 0 {
				word = append(word, edges[0])
				if !d.dictionary.IsPrefix(word) {
					// backtrack word
					word = word[:len(word)-1]
					continue
				}
				if d.dictionary.IsWord(word) {
					answers = append(answers, string(word))
				}
			}
		}
	}
	return nil, nil
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

func getEdges(x, y int, board [][]rune) []rune {
	const minX, minY = 0, 0
	maxY := len(board)
	maxX := len(board[0])

	posX := x - 1
	posY := y - 1

	output := make([]rune, 0)

	for posY < maxY && posY <= y+1 {
		if posY < minY {
			posY++
		}
		for posX < maxX && posX <= x+1 {
			if posX < minX {
				posX++
				continue
			}
			if !(posY == y && posX == x) {
				output = append(output, board[posY][posX])
			}
			posX++
		}
		posY++
		posX = x - 1
	}

	return output
}
