package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type KthLargest struct {
	k  int
	pq *priorityqueue.Queue
}

func Constructor(k int, nums []int) KthLargest {
	pq := priorityqueue.NewWith(func(a, b any) int {
		return cmp.Compare(a.(int), b.(int))
	})
	for _, n := range nums {
		pq.Enqueue(n)
		if pq.Size() > k {
			pq.Dequeue()
		}
	}
	return KthLargest{
		k: k, pq: pq,
	}
}

func (this *KthLargest) Add(val int) int {
	for this.pq.Enqueue(val); this.pq.Size() > this.k; this.pq.Dequeue() {
	}
	res, _ := this.pq.Peek()
	return res.(int)
}
