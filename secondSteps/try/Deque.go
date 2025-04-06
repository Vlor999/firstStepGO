package try

import (
	"fmt"
	"sync"
)

type Node struct {
	value []int
	prev  *Node
	next  *Node
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

// PushFront ajoute un élément en tête
func (d *Deque) PushFront(value []int) {
	node := &Node{value: value}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.size++
}

// PushBack ajoute un élément en queue
func (d *Deque) PushBack(value []int) {
	node := &Node{value: value}
	if d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.size++
}

// PopFront enlève et retourne l'élément en tête
func (d *Deque) PopFront() []int {
	if d.head == nil {
		return nil
	}
	val := d.head.value
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
	return val
}

// PopBack enlève et retourne l'élément en queue
func (d *Deque) PopBack() []int {
	if d.tail == nil {
		return nil
	}
	val := d.tail.value
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
	return val
}

// GetFront retourne la valeur en tête
func (d *Deque) GetFront() *Node {
	return d.head
}

// GetBack retourne la valeur en queue
func (d *Deque) GetBack() *Node {
	return d.tail
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetValue() []int {
	return n.value
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
	*result = *result && (dist >= radius*radius)
	mu.Unlock()
}


func (d *Deque) Verify(radius int) bool {
	if d.head == nil {
		return true
	}

	length := 0
	for node := d.head; node != nil; node = node.next {
		length++
	}

	if length <= 2*radius {
		return true
	}

	// On récupère le head
	head := d.head
	result := true
	var wg sync.WaitGroup
	var mu sync.Mutex

	// On parcourt les éléments à partir de l'index 2*radius + 1
	index := 0
	for node := d.head; node != nil; node = node.next {
		if index > 2*radius {
			wg.Add(1)
			go localVerify(head.value, node.value, radius, &result, &wg, &mu)
		}
		index++
	}

	wg.Wait()
	return result
}


// Size retourne le nombre d’éléments
func (d *Deque) Size() int {
	return d.size
}

// Print affiche tous les éléments de la deque
func (d *Deque) Print() {
	curr := d.head
	for curr != nil {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
	fmt.Println()
}
