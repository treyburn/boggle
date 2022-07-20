package solver

type Node struct {
	char rune
	x    int
	y    int
}

func buildGraph(input [][]rune) [][]*Node {
	graph := make([][]*Node, 0)
	for yIdx, row := range input {
		itemRow := make([]*Node, 0)
		for xIdx, r := range row {
			n := &Node{
				char: r,
				x:    xIdx,
				y:    yIdx,
			}
			itemRow = append(itemRow, n)
		}
		graph = append(graph, itemRow)
	}
	return graph
}

func getEdges(node *Node, graph [][]*Node) []*Node {
	const minX, minY = 0, 0
	maxY := len(graph)
	maxX := len(graph[0])

	posX := node.x - 1
	posY := node.y - 1

	output := make([]*Node, 0)

	for posY < maxY && posY <= node.y+1 {
		if posY < minY {
			posY++
		}
		for posX < maxX && posX <= node.x+1 {
			if posX < minX {
				posX++
				continue
			}
			if !(posY == node.y && posX == node.x) {
				output = append(output, graph[posY][posX])
			}
			posX++
		}
		posY++
		posX = node.x - 1
	}

	return output
}
