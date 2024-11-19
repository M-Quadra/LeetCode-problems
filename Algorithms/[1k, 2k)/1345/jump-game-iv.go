type note struct {
	to   int
	next int
}

func minJumps(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	head := func() []int { // 预处理
		uniMap := make(map[int]int, len(arr))
		uniMap[arr[0]] = 0
		arr[0] = -1

		offset := 0
		for i := 1; i < len(arr); i++ {
			n, ok := uniMap[arr[i]]
			if !ok {
				n = len(uniMap)
				uniMap[arr[i]] = n
			}
			arr[i] = n
			if i > 1 && arr[i-2] == arr[i] && arr[i-1] == arr[i] {
				offset++
				continue
			}

			arr[i-offset] = n
		}

		arr = arr[:len(arr)-offset]
		arr[0] = 0

		head := make([]int, len(uniMap))
		for i := range head {
			head[i] = -1
		}
		return head
	}()

	table := make([]note, 0, len(arr))
	for i, v := range arr {
		n := note{
			to:   i,
			next: head[v],
		}
		head[v] = len(table)
		table = append(table, n)
	}

	used := make([]bool, len(arr))
	queue := make([]int, len(arr)-len(head)+2)
	i, j := 0, 1
	push := func(v int) {
		if used[v] {
			return
		}
		used[v] = true
		queue[j] = v
		j++
		j %= len(queue)
	}

	step := make([]int, len(arr))
	for i := 1; i < len(step); i++ {
		step[i] = 0x0FFFFFFF
	}

	for i != j {
		now := queue[i]
		i++
		i %= len(queue)
		used[now] = false
		nStep := step[now] + 1

		for i := head[arr[now]]; i != -1; i = table[i].next {
			if v := table[i].to; nStep < step[v] {
				step[v] = nStep
				push(v)
			}
		}

		if now-1 > 0 {
			if v := now - 1; nStep < step[v] {
				step[v] = nStep
				push(v)
			}
		}
		if now+1 < len(arr) {
			if v := now + 1; nStep < step[v] {
				step[v] = nStep
				push(v)
			}
		}
	}
	return step[len(step)-1]
}