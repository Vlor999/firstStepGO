package try

import (
	"fmt"
	"sync"
)

type Deque struct {
	Data [][]int
}

func (d *Deque) GetFront() []int {
	if len(d.Data) == 0 {
		return nil
	}
	return d.Data[0]
}

func (d *Deque) GetQueue() []int {
	if len(d.Data) == 0 {
		return nil
	}
	return d.Data[len(d.Data)-1]
}

func (d *Deque) PushFront(value []int) {
	d.Data = append([][]int{value}, d.Data...)
}

func (d *Deque) PushBack(value []int) {
	d.Data = append(d.Data, value)
}

func (d *Deque) PopFront() ([]int, bool) {
	if len(d.Data) == 0 {
		return []int{0}, false
	}
	front := d.Data[0]
	d.Data = d.Data[1:]
	return front, true
}

func (d *Deque) PopBack() ([]int, bool) {
	if len(d.Data) == 0 {
		return []int{0}, false
	}
	back := d.Data[len(d.Data)-1]
	d.Data = d.Data[:len(d.Data)-1]
	return back, true
}

func localVerify(pos1 []int, pos2 []int, radius int, result *bool, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	if !*result {
		mu.Unlock()
		return
	}
	mu.Unlock()

	dist := (pos1[0]-pos2[0])*(pos1[0]-pos2[0]) + (pos1[1]-pos2[1])*(pos1[1]-pos2[1])
	mu.Lock()
	*result = dist >= radius*radius
	mu.Unlock()
}

func (d *Deque) Verify(radius int) bool {
	if d.Size() <= 2*radius {
		return true
	}

	head := d.GetFront()
	data := d.Data[2*radius+1:]
	result := true

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, current := range data {
		wg.Add(1)
		go localVerify(head, current, radius, &result, &wg, &mu)
	}

	wg.Wait()
	return result
}


func (d *Deque) Print() {
	fmt.Println(d.Data)
}

func (d *Deque) Size() int {
	return len(d.Data)
}
