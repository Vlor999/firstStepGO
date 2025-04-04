package try

import "fmt"

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

func (d *Deque) Verify(radius int) bool {
	if d.Size() <= 2*radius {
		return true
	}

	radiusSquared := radius * radius
	head := d.GetFront()
	headX, headY := head[0], head[1]

	for _, current := range d.Data[2*radius+1:] {
		dx, dy := headX-current[0], headY-current[1]
		if dx*dx+dy*dy < radiusSquared {
			return false
		}
	}
	return true
}

func (d *Deque) Print() {
	fmt.Println(d.Data)
}

func (d *Deque) Size() int {
	return len(d.Data)
}
