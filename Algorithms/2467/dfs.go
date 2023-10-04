func buildTree(edges [][]int, bob int32, amount []int) {
	stepA := make([]int32, len(amount))
	parents := make([]int32, 0, len(amount))

	heads := make([]int32, 0, len(amount))
	allEdges := make([][2]int32, 0, len(edges)*2)
	for range amount {
		parents = append(parents, -1)
		heads = append(heads, -1)
	}

	for _, edge := range edges {
		u, v := int32(edge[0]), int32(edge[1])
		allEdges = append(allEdges, [2]int32{v, heads[u]})
		heads[u] = int32(len(allEdges) - 1)
		allEdges = append(allEdges, [2]int32{u, heads[v]})
		heads[v] = int32(len(allEdges) - 1)
	}

	iE := 0
	var dfs func(from int32)
	dfs = func(from int32) {
		for i := heads[from]; i > -1; i = allEdges[i][1] {
			to := allEdges[i][0]
			if parents[to] != -1 || to == 0 {
				continue
			}
			parents[to] = from
			stepA[to] = stepA[from] + 1
			edges[iE][0], edges[iE][1] = int(from), int(to)
			iE++
			dfs(to)
		}
	}
	dfs(0)

	for i, s := bob, int32(0); i > -1; i, s = int32(parents[i]), s+1 {
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

	for i, edge := range edges {
		u, v := edge[0], edge[1]
		amount[v] += amount[u]
		if (i == len(edges)-1 || edges[i+1][0] != v) && amount[v] > opt {
			opt = amount[v]
		}
	}
	return opt
}
