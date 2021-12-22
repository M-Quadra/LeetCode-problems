func sumGame(num string) bool {
	sum, qm := []int{0, 0}, []int{0, 0}
	for i, v := range num {
		i = i / (len(num) >> 1)
		if v == '?' {
			qm[i]++
			continue
		}
		sum[i] += int(v - '0')
	}
	sum[0] -= sum[1]
	qm[0] -= qm[1]
	if sum[0] == 0 {
		return qm[0] != 0
	}

	if sum[0] < 0 {
		sum[0], qm[0] = -sum[0], -qm[0]
	}
	qm[0] = -qm[0]
	if qm[0] <= 0 {
		return true
	}

	qm[1] = qm[0] >> 1
	qm[0] -= qm[1]
	return qm[0]*9 > sum[0] || qm[1]*9 < sum[0]
}