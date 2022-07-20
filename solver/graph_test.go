package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildGraph(t *testing.T) {
	rb := [][]rune{
		{1, 2, 3},
	}

	graph := buildGraph(rb)
	assert.Equal(t, 1, len(graph))
	assert.Equal(t, 3, len(graph[0]))

	assert.Equal(t, rune(1), graph[0][0].char)
	assert.Equal(t, 0, graph[0][0].x)
	assert.Equal(t, 0, graph[0][0].y)
	assert.Equal(t, rune(2), graph[0][1].char)
	assert.Equal(t, 1, graph[0][1].x)
	assert.Equal(t, 0, graph[0][1].y)
	assert.Equal(t, rune(3), graph[0][2].char)
	assert.Equal(t, 2, graph[0][2].x)
	assert.Equal(t, 0, graph[0][2].y)
}

func Test_getEdges(t *testing.T) {
	type testCase struct {
		name      string
		board     [][]*Node
		node      *Node
		wantEdges []*Node
	}

	n00 := &Node{
		x: 0,
		y: 0,
	}

	n10 := &Node{
		x: 1,
		y: 0,
	}

	n20 := &Node{
		x: 2,
		y: 0,
	}

	n01 := &Node{
		x: 0,
		y: 1,
	}

	n11 := &Node{
		x: 1,
		y: 1,
	}

	n21 := &Node{
		x: 2,
		y: 1,
	}

	n02 := &Node{
		x: 0,
		y: 2,
	}

	n12 := &Node{
		x: 1,
		y: 2,
	}

	n22 := &Node{
		x: 2,
		y: 2,
	}

	testBoard := [][]*Node{
		{n00, n10, n20},
		{n01, n11, n21},
		{n02, n12, n22},
	}

	var tests = []testCase{
		{"3x3 - middle", testBoard, n11, []*Node{n00, n10, n20, n01, n21, n02, n12, n22}},
		{"3x3 - upper center", testBoard, n10, []*Node{n00, n20, n01, n11, n21}},
		{"3x3 - upper left", testBoard, n00, []*Node{n10, n01, n11}},
		{"3x3 - lower right", testBoard, n22, []*Node{n11, n21, n12}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := getEdges(tc.node, tc.board)

			assert.Equal(t, tc.wantEdges, got)
		})
	}
}
