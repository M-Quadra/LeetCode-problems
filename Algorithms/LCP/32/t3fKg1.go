func processTasks(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})

	i := 0
	for _, v := range tasks {
		if i <= 0 {
			v[0] = v[1] - v[2] + 1
			tasks[i] = v
			i++
			continue
		}

		last := tasks[i-1]
		if last[1] < v[0] {
			v[0] = v[1] - v[2] + 1
			v[2] += last[2]
			tasks[i] = v
			i++
			continue
		}

		j := 0
		for count, half, middle := i, 0, 0; count > 0; {
			half = count >> 1
			middle = j + half
			if tasks[middle][1] < v[0] {
				j = middle + 1
				count -= half + 1
				continue
			}

			count = half
		}

		p, t := last[2], tasks[j]
		if j > 0 {
			p -= tasks[j-1][2]
		}
		if t[0] <= v[0] && v[0] <= t[1] {
			p -= v[0] - t[0]
		}

		v[2] -= p
		if v[2] <= 0 {
			continue
		}

		v[0] = v[1] - v[2] + 1
		v[2] += last[2]
		tasks[i] = v
		// 区间压缩
		for ; i > 0 && tasks[i][0] <= tasks[i-1][1]; i-- {
			copy(tasks[i-1][1:], tasks[i][1:])
			tasks[i-1][0] = tasks[i-1][1] - tasks[i-1][2] + 1
			if i-1 > 0 {
				tasks[i-1][0] += tasks[i-2][2]
			}
		}
		i++
	}

	if i <= 0 {
		return 0
	}
	return tasks[i-1][2]
}