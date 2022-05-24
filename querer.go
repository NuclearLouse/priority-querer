package querer

import (
	"container/heap"
)

/*
package main

import (
	"container/heap"

	querer "utilits/priority-querer"
)

var jobQueue querer.PriorityQueue

type item struct {
	Order int
}

func main() {
	jobQueue = make(querer.PriorityQueue, 0)
	heap.Init(&jobQueue)
	var items []*item
	items = []*item{
		{Order: 1},
		{Order: 5},
		{Order: 0},
		{Order: -1},
		{Order: 4},
		{Order: 3},
		{Order: 1},
	}
	for _, i := range items {
		heap.Push(&jobQueue, &querer.Item{Value: i, Priority: i.Order})
	}

	for queue.Len() > 0 {
		e := heap.Pop(jobQueue).(*querer.Item)
		i := := e.Value.(*item)

		Do work with item ....

	}

}
*/

// An Item is something we manage in a priority queue.
type Item struct {
	Value    interface{} // The value of the item; arbitrary.
	Priority int         // The priority of the item in the queue.
	index    int         // The index of the item in the heap. The index is needed by update and is maintained by the heap.Interface methods.
}

type PriorityQueue []*Item

type any = interface{}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest (=0), not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value any, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
