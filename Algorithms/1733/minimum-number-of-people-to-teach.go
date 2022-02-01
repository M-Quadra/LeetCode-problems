func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	lanSet := make([][]bool, n)
	for i := range lanSet {
		lanSet[i] = make([]bool, len(languages))
	}

	for i, langs := range languages {
		for _, v := range langs {
			lanSet[v-1][i] = true
		}
	}

	for i := 0; i < len(friendships); i++ {
		friendships[i][0]--
		friendships[i][1]--

		a, b := friendships[i][0], friendships[i][1]
		ok := false
		for j := range lanSet {
			if lanSet[j][a] == true && lanSet[j][b] == true {
				ok = true
				break
			}
		}

		if ok {
			friendships[len(friendships)-1], friendships[i] = friendships[i], friendships[len(friendships)-1]
			friendships = friendships[:len(friendships)-1]
			i--
		}
	}

	opt := len(languages)
	for i := range lanSet {
		cnt := 0
		for _, edge := range friendships {
			for _, v := range edge {
				if !lanSet[i][v] {
					lanSet[i][v] = true
					cnt++
				}
			}
			if cnt >= opt {
				break
			}
		}
		opt = min(opt, cnt)
	}

	return opt
}