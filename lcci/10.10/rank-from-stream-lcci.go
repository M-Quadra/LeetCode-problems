type StreamRank struct {
	tree [50005]int
}

func Constructor() StreamRank {
	return StreamRank{}
}

func (this *StreamRank) Track(x int) {
	for x++; x < len(this.tree); x += x & -x {
		this.tree[x]++
	}
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	opt := 0
	for x++; x > 0; x -= x & -x {
		opt += this.tree[x]
	}
	return opt
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */