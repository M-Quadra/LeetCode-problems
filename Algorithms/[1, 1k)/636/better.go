func exclusiveTime(n int, logs []string) []int {
	res, s := make([]int, n), 0
	var stack []int
	for _, log := range logs {
		t := strings.Index(log, ":")
		i, _ := strconv.Atoi(log[:t])
		log = log[t+1:]
		t, _ = strconv.Atoi(log[strings.Index(log, ":")+1:])
		if t -= s; log[:3] == "sta" {
			stack = append(stack, t)
			continue
		}
		v := t - stack[len(stack)-1] + 1
		stack = stack[:len(stack)-1]
		res[i] += v
		s += v
	}
	return res
}
