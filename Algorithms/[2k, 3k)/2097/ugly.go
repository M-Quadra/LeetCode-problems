func validArrangement(pairs [][]int) [][]int {
	preMapA := make(map[int][][]int, len(pairs)/4)
	preMapB := make(map[int][][]int, len(pairs)/8)

	prev := make([]int, len(pairs))
	next := make([]int, len(pairs))
	for i, v := range pairs {
		prev[i] = i
		next[i] = i
		preMapA[v[0]] = append(preMapA[v[0]], []int{i, i})
	}

	cyclePreDic := map[int][]int{}

	for len(preMapA) > 0 {
		k, v := 0, [][]int(nil)
		for nK, nV := range preMapA {
			k, v = nK, nV
			break
		}

		edge := v[len(v)-1]
		v = v[:len(v)-1]
		if len(v) <= 0 {
			delete(preMapA, k)
		} else {
			preMapA[k] = v
		}

		for {
			if pre, ok := cyclePreDic[pairs[edge[1]][1]]; ok { // 窃取环
				delete(cyclePreDic, pairs[edge[1]][1])
				prev[pre[0]] = edge[1]
				next[edge[1]] = pre[0]
				edge[1] = pre[1]
				continue
			}

			if pre, ok := preMapB[pairs[edge[1]][1]]; ok { // 窃取边
				e := pre[len(pre)-1]
				pre = pre[:len(pre)-1]
				if len(pre) <= 0 {
					delete(preMapB, pairs[edge[1]][1])
				} else {
					preMapB[pairs[edge[1]][1]] = pre
				}

				prev[e[0]] = edge[1]
				next[edge[1]] = e[0]
				edge[1] = e[1]
				continue
			}
			if pre, ok := preMapA[pairs[edge[1]][1]]; ok {
				e := pre[len(pre)-1]
				pre = pre[:len(pre)-1]
				if len(pre) <= 0 {
					delete(preMapA, pairs[edge[1]][1])
				} else {
					preMapA[pairs[edge[1]][1]] = pre
				}

				prev[e[0]] = edge[1]
				next[edge[1]] = e[0]
				edge[1] = e[1]
				continue
			}

			// 无法拼接
			if pairs[edge[0]][0] == pairs[edge[1]][1] { // 环
				k := pairs[edge[0]][0]
				if pre, ok := cyclePreDic[k]; ok {
					next[pre[1]] = edge[0]
					pre[1] = edge[1]
					cyclePreDic[k] = pre
				} else {
					cyclePreDic[k] = edge
				}
			} else { // 缓存 边
				k := pairs[edge[0]][0]
				preMapB[k] = append(preMapB[k], edge)
			}
			break
		}
	}

	// 构建索引
	inCycleMap := map[int][]int{}

	for k, v := range cyclePreDic {
		// delete(cyclePreDic, k)

		for i := v[0]; ; i = next[i] {
			cycIndex, ok := inCycleMap[pairs[i][1]]
			if ok {
				if cycIndex[0] == k {
					goto buildIndex
				}

				delete(inCycleMap, pairs[i][1])
				if cycle, ok := cyclePreDic[cycIndex[0]]; ok {
					delete(cyclePreDic, cycIndex[0])

					next[cycle[1]] = cycle[0]
					prev[cycle[0]] = cycle[1]

					nextC := prev[cycIndex[1]]
					nextI := next[i]

					next[i] = cycIndex[1]
					prev[cycIndex[1]] = i
					next[nextC] = nextI
					prev[nextI] = nextC

					// 不会触发终点调整, 因为在首项处理过了
				}
			}

		buildIndex:
			if cycIndex, ok := inCycleMap[pairs[i][0]]; ok {
				if cycIndex[0] == k { // 没必要刷新
					goto ed
				}

				// 环起点与某环起点重合, 此处直接合并2个环
				if cycle, ok := cyclePreDic[cycIndex[0]]; ok {
					delete(cyclePreDic, cycIndex[0])

					next[cycle[1]] = cycle[0]
					prev[cycle[0]] = cycle[1]

					nextC := prev[cycIndex[1]]

					next[v[1]] = cycIndex[1]
					prev[cycIndex[1]] = v[1]

					next[nextC] = nextC
					v[1] = nextC
					goto ed
				}

				cycIndex[0], cycIndex[1] = k, i
				inCycleMap[pairs[i][0]] = cycIndex
			} else {
				inCycleMap[pairs[i][0]] = []int{k, i}
			}
		ed:
			if i == v[1] {
				break
			}
		}

	}

	opt := make([][]int, 0, len(pairs))
	first := true
	for _, lines := range preMapB {
		line := lines[0]
		for i := line[0]; ; i = next[i] {
			if cycIndex, ok := inCycleMap[pairs[i][0]]; ok && first {
				first = false
				if cycle, ok := cyclePreDic[cycIndex[0]]; ok {
					delete(cyclePreDic, cycIndex[0])

					next[cycle[1]] = cycle[0]
					prev[cycle[0]] = cycle[1]

					nextC := prev[cycIndex[1]]
					next[nextC] = line[0]
					prev[line[0]] = nextC

					line[0] = cycIndex[1]
					prev[line[0]] = line[0]
					i = line[0]
				}
			}

			opt = append(opt, pairs[i])

			if cycIndex, ok := inCycleMap[pairs[i][1]]; ok {
				if cycle, ok := cyclePreDic[cycIndex[0]]; ok {
					delete(cyclePreDic, cycIndex[0])

					next[cycle[1]] = cycle[0]
					prev[cycle[0]] = cycle[1]

					nextC := prev[cycIndex[1]]
					if i == line[1] {
						next[i] = cycIndex[1]
						prev[cycIndex[1]] = i

						line[1] = nextC
						next[nextC] = nextC
					} else {
						nextI := next[i]

						next[i] = cycIndex[1]
						prev[cycIndex[1]] = i

						next[nextC] = nextI
						prev[nextI] = nextC
					}
				}
			}

			if i == line[1] {
				break
			}

		}

		break
	}

	for _, v := range cyclePreDic {
		for i := v[0]; ; i = next[i] {
			opt = append(opt, pairs[i])
			if i == v[1] {
				break
			}
		}
	}
	return opt
}
