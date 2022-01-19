func sequenceReconstruction(org []int, seqs [][]int) bool {
	{ // 预处理
		mp := make([]int, len(org)+1)
		for i, v := range org {
			mp[v] = i + 1
			org[i] = i + 1
		}
		for _, seq := range seqs {
			for i, v := range seq {
				if v <= 0 || v > len(org) {
					return false
				}

				seq[i] = mp[v]
			}
		}
	}

	if len(org) == 1 {
		for _, seq := range seqs {
			if len(seq) != 1 {
				return false
			}

			return seq[0] == 1
		}
		return false
	}

	for _, seq := range seqs {
		for i, v := range seq {
			if i+1 >= len(seq) {
				continue
			}

			next := seq[i+1]
			if next <= v {
				return false
			}

			if next == v+1 {
				org[v-1] = 0
			}
		}
	}

	for _, v := range org[:len(org)-1] {
		if v != 0 {
			return false
		}
	}

	return true
}