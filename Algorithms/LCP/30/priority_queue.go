package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func magicTower(nums []int) int {
	pq := priorityqueue.NewWith(func(a, b any) int {
		return cmp.Compare(a.(int), b.(int))
	})
	hp, cnt, atk := 1, 0, 0
	for _, num := range nums {
		hp += num
		if num < 0 {
			pq.Enqueue(num)
		}
		for ; !pq.Empty() && hp <= 0; cnt++ {
			v, _ := pq.Dequeue()
			hp -= v.(int)
			atk += v.(int)
		}
	}
	hp += atk
	if hp > 0 {
		return cnt
	}
	return -1
}
