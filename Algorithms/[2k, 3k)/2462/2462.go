import (
	"math"
	"slices"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func totalCost(costs []int, k, candidates int) int64 {
	if len(costs) <= candidates*2 {
		slices.Sort(costs)
		ans := int64(0)
		for k--; k >= 0; k-- {
			ans += int64(costs[k])
		}
		return ans
	}

	lQue := pq.NewWith(utils.IntComparator)
	rQue := pq.NewWith(utils.IntComparator)
	l, r := 0, len(costs)-1
	for i := 0; i < candidates; i++ {
		lQue.Enqueue(costs[l])
		rQue.Enqueue(costs[r])
		l, r = l+1, r-1
	}

	ans := int64(0)
	for ; k > 0; k-- {
		lV, rV := math.MaxInt, math.MaxInt
		if v, ok := lQue.Peek(); ok {
			lV = v.(int)
		}
		if v, ok := rQue.Peek(); ok {
			rV = v.(int)
		}
		ans += int64(min(lV, rV))

		if lV <= rV {
			if lQue.Dequeue(); l <= r {
				lQue.Enqueue(costs[l])
				l++
			}
		} else if rQue.Dequeue(); l <= r {
			rQue.Enqueue(costs[r])
			r--
		}
	}
	return ans
}
