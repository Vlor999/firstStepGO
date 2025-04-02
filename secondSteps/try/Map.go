package try

func UpdateMap(direction [2]int, dequePosition *Deque) *Deque {
    if len(dequePosition.Data) == 0 {
        dequePosition.PushBack([]int{400, 300})
        return dequePosition
    }
    
    x := direction[0]
    y := direction[1]
    head := dequePosition.GetFront()
    if head == nil {
        dequePosition.PushBack([]int{400, 300})
        return dequePosition
    }
    
    dequePosition.PushFront([]int{head[0] + x, head[1] + y})
    _, _ = dequePosition.PopBack()
    return dequePosition
}