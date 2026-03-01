package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/trees/avltree"
)

var _sums []int64

type seg [2]int // [st, ed)

func (s seg) sum() int64 {
	return _sums[s[1]] - _sums[s[0]]
}

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	_sums = make([]int64, len(nums)+1)
	for i, n := range nums {
		_sums[i+1] = _sums[i] + int64(n)
	}
	segPq := priorityqueue.NewWith(func(a, b any) int {
		return cmp.Compare(b.(seg).sum(), a.(seg).sum())
	})
	segPq.Enqueue(seg{0, len(nums)})

	avl := avltree.NewWithIntComparator()
	res := make([]int64, 0, len(nums))
	for _, q := range removeQueries {
		avl.Put(q, nil)

	loop:
		var top seg
		if j, ok := segPq.Peek(); ok {
			top = j.(seg)
		}

		for n := avl.Root; n != nil; {
			q = n.Key.(int)
			switch {
			case top[1] < q:
				n = n.Children[0]
			case q < top[0]:
				n = n.Children[1]
			default:
				n = nil
			}
		}
		if top[0] <= q && q < top[1] {
			avl.Remove(q)
			if segPq.Dequeue(); top[0] < q {
				segPq.Enqueue(seg{top[0], q})
			}
			if q+1 < top[1] {
				segPq.Enqueue(seg{q + 1, top[1]})
			}
			goto loop
		}
		res = append(res, top.sum())
	}
	return res
}
