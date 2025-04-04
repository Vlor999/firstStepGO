package try

func UpdateMap(direction [2]int, dequePosition *Deque, mustDelete bool) {
    if len(dequePosition.Data) == 0 {
        dequePosition.PushBack([]int{400, 300})
    }
    
    x := direction[0]
    y := direction[1]
    head := dequePosition.GetFront()
    if head == nil {
        dequePosition.PushBack([]int{400, 300})
    }
    
    dequePosition.PushFront([]int{head[0] + x, head[1] + y})
    if mustDelete{
        _, _ = dequePosition.PopBack()
    }
}

func IsInBounds(head []int, maxX int, maxY int) bool {
    x := head[0]
    y := head[1]
    return 0 <= x && x < maxX && 0 <= y && y < maxY
}

func HandleSnakeApple(dequePosition *Deque, randomPoint [2]int, radius int) int {
	head := dequePosition.GetFront()
	x_snake := head[0]
	y_snake := head[1]
	x_random := randomPoint[0]
	y_random := randomPoint[1]
	if (x_snake-x_random)*(x_snake-x_random)+(y_snake-y_random)*(y_snake-y_random) <= radius*radius {
		return radius
	}
	return 0
}
