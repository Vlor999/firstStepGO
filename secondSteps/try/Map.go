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

func CheckSelfCollision(deque *Deque) bool {
	if len(deque.Data) <= 1 {
		return false
	}

	head := deque.GetFront()

	for i := 1; i < len(deque.Data); i++ {
		segment := deque.Data[i]
		// Si les coordonnées sont très proches, il y a collision
		if abs(head[0]-segment[0]) < 5 && abs(head[1]-segment[1]) < 5 {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
