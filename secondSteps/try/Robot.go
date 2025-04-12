package try

import (
	"fmt"
)

type Robot struct {
	Snake Deque
	road  Deque
	apple [2]int
}

func (d *Robot) Print() {
	d.Snake.Print()
	fmt.Print(" is going to  : ")
	fmt.Print(d.apple)
}

func (d *Robot) Println() {
	d.Print()
	fmt.Println()
}

func heuristic(pos1 [2]int, pos2 [2]int) int {
	return (pos1[0]-pos2[0])*(pos1[0]-pos2[0]) + (pos1[1]-pos2[1])*(pos1[1]-pos2[1])
}

// SetSnake initializes the robot's internal snake with the current game snake position
func (d *Robot) SetSnake(gameSnake *Deque) {
	d.Snake = Deque{}
	if gameSnake != nil && gameSnake.head != nil {
		// Copy the head position
		d.Snake.PushBack([]int{gameSnake.head.value[0], gameSnake.head.value[1]})
	} else {
		// Default position if no snake exists
		d.Snake.PushBack([]int{400, 300})
	}
}

func (d *Robot) SetPath() {
	// Make sure snake is initialized
	if d.Snake.head == nil {
		d.Snake.PushBack([]int{400, 300})
	}

	start := [2]int{d.Snake.head.value[0], d.Snake.head.value[1]}
	goal := d.apple

	openSet := make(map[[2]int]bool)
	openSet[start] = true

	cameFrom := make(map[[2]int][2]int)

	gScore := make(map[[2]int]int)
	gScore[start] = 0

	fScore := make(map[[2]int]int)
	fScore[start] = heuristic(start, goal)

	for len(openSet) > 0 {
		var current [2]int
		lowestFScore := int(^uint(0) >> 1)
		for node := range openSet {
			if fScore[node] < lowestFScore {
				lowestFScore = fScore[node]
				current = node
			}
		}

		if current == goal {
			d.road = reconstructPath(cameFrom, current)
			return
		}

		delete(openSet, current)

		for _, neighbor := range getNeighbors(current) {
			tentativeGScore := gScore[current] + 1

			if g, exists := gScore[neighbor]; !exists || tentativeGScore < g {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + heuristic(neighbor, goal)

				if !openSet[neighbor] {
					openSet[neighbor] = true
				}
			}
		}
	}

	// If no path found, create a simple path toward the apple
	d.road = Deque{}
	if d.Snake.head != nil && d.Snake.head.value != nil {
		dx := d.apple[0] - d.Snake.head.value[0]
		dy := d.apple[1] - d.Snake.head.value[1]

		// Move horizontally first
		if dx != 0 {
			dir := 1
			if dx < 0 {
				dir = -1
			}
			d.road.PushBack([]int{d.Snake.head.value[0] + dir, d.Snake.head.value[1]})
		}

		// Then vertically
		if dy != 0 {
			dir := 1
			if dy < 0 {
				dir = -1
			}
			d.road.PushBack([]int{d.apple[0], d.Snake.head.value[1] + dir})
		}
	}
}

func reconstructPath(cameFrom map[[2]int][2]int, current [2]int) Deque {
	path := Deque{}
	for {
		prev, exists := cameFrom[current]
		if !exists {
			break
		}
		newPrev := []int{prev[0], prev[1]}
		path.PushFront(newPrev)
		current = prev
	}
	return path
}

func getNeighbors(node [2]int) [][2]int {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	neighbors := make([][2]int, 0, len(directions))
	for _, dir := range directions {
		neighbor := [2]int{node[0] + dir[0], node[1] + dir[1]}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func (d *Robot) SetApplePosition(position [2]int) {
	d.apple = position
}

func (d *Robot) GetNextDirection() [2]int {
	// Default direction if there's no valid path
	defaultDirection := [2]int{1, 0}

	// Check if snake is initialized
	if d.Snake.head == nil || d.Snake.head.value == nil {
		return defaultDirection
	}

	// Check if we have a path
	if d.road.size == 0 {
		// Recalculate path
		d.SetPath()

		// Still no path? Return default
		if d.road.size == 0 {
			return defaultDirection
		}
	}

	start := d.Snake.head.value
	next := d.road.PopFront()

	// If next is nil, return default direction
	if next == nil {
		return defaultDirection
	}

	return [2]int{next[0] - start[0], next[1] - start[1]}
}
