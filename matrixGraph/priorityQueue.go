package matrixGraph

import (
	"container/heap"
)

type space struct { // dis from x to end , x have know
	end   int
	cost  int
	index int
}
type priorityQueue []*space

func (this priorityQueue) Len() int {
	return len(this)
}
func (this priorityQueue) Less(i, j int) bool {
	return this[i].cost < this[j].cost //pop the lowest
}

func (this priorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *priorityQueue) Push(x interface{}) {
	n := len(*this)
	node := x.(*space)
	node.index = n

	*this = append(*this, node)
}

func (this *priorityQueue) Pop() interface{} {
	old := *this
	n := len(old)
	node := old[n-1]
	node.index = -1 //for safety
	*this = old[0 : n-1]
	return node
}

func (this *priorityQueue) update(node *space, end int, cost int) {
	heap.Remove(this, node.index)
	node.end = end
	node.cost = cost
	heap.Push(this, node)

}
