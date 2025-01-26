type edge struct {
	to   int
	next int
}

func validArrangement(pairs [][]int) [][]int {
	encodingMap := map[int]int{}
	for i := range pairs {
		for j := range pairs[i] {
			v, ok := encodingMap[pairs[i][j]]
			if !ok {
				v = len(encodingMap)
				encodingMap[pairs[i][j]] = v
			}
			pairs[i][j] = v
		}
	}

	func() {
		head := make([]int, len(encodingMap))
		for i := range head {
			head[i] = -1
		}

		table := make([]edge, 0, len(pairs))
		st := 0
		func() {
			indegree := make([]int, len(encodingMap))
			outdegree := make([]int, len(encodingMap))

			for _, v := range pairs {
				outdegree[v[0]]++
				indegree[v[1]]++

				e := edge{to: v[1], next: head[v[0]]}
				head[v[0]] = len(table)
				table = append(table, e)
			}

			for i := range indegree {
				if indegree[i] < outdegree[i] {
					st = i
					break
				}
			}
		}()

		pairs[0][0], pairs[0][1] = st, table[head[st]].to
		head[st] = table[head[st]].next
		for i, j := 0, len(pairs)-1; j >= 0; {
			now := pairs[i]
			if head[now[1]] == -1 {
				pairs[j][0], pairs[j][1] = pairs[i][0], pairs[i][1]
				j--

				if head[now[0]] != -1 {
					pairs[i][0], pairs[i][1] = now[0], table[head[now[0]]].to
					head[now[0]] = table[head[now[0]]].next
				} else {
					i--
				}
				continue
			}

			i++
			pairs[i][0], pairs[i][1] = now[1], table[head[now[1]]].to
			head[now[1]] = table[head[now[1]]].next
		}
	}()

	decodingArr := make([]int, len(encodingMap))
	for k, v := range encodingMap {
		decodingArr[v] = k
	}
	for i := range pairs {
		for j := range pairs[i] {
			pairs[i][j] = decodingArr[pairs[i][j]]
		}
	}
	return pairs
}