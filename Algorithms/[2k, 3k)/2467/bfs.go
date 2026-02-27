var (
	_parents []int32
	_que     []int32

	_heads []int32
	_edges [][2]int32
)

func buildTree(edges [][]int, bob int32, amount []int) {
	stepA := make([]int32, 0, len(amount))
	_parents = make([]int32, 0, len(amount))

	_heads = make([]int32, 0, len(amount))
	_edges = make([][2]int32, 0, len(edges)*2)
	for i := 0; i < len(amount); i++ {
		stepA = append(stepA, -1)
		_parents = append(_parents, -1)
		_heads = append(_heads, -1)
	}

	for _, edge := range edges {
		u, v := int32(edge[0]), int32(edge[1])
		_edges = append(_edges, [2]int32{v, _heads[u]})
		_heads[u] = int32(len(_edges) - 1)
		_edges = append(_edges, [2]int32{u, _heads[v]})
		_heads[v] = int32(len(_edges) - 1)
	}

	stepA[0] = 0
	_que = make([]int32, 1, len(amount))
	for i := 0; i < len(_que); i++ {
		from := _que[i]

		for j := _heads[from]; j > -1; j = _edges[j][1] {
			to := _edges[j][0]
			if _parents[to] != -1 || to == 0 {
				continue
			}
			_parents[to] = from
			stepA[to] = stepA[from] + 1
			_que = append(_que, to)
		}
	}

	for i, s := bob, int32(0); i > -1; i, s = int32(_parents[i]), s+1 {
		if stepA[i] == s {
			amount[i] /= 2
		} else if stepA[i] > s {
			amount[i] = 0
		}
	}
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	buildTree(edges, int32(bob), amount)
	opt := math.MinInt

	for _, from := range _que {
		cnt := 0
		for i := _heads[from]; i > -1; i = _edges[i][1] {
			if to := _edges[i][0]; to != from && to != _parents[from] {
				cnt++

				amount[to] += amount[from]
			}
		}
		if cnt <= 0 && amount[from] > opt {
			opt = amount[from]
		}
	}
	return opt
}
