func distributeCandies(candies []int) int {
	mp := map[int]struct{}{}
	for _, c := range candies {
		mp[c] = struct{}{}
	}

	cnt := len(candies) / 2
	if len(mp) < cnt {
		return len(mp)
	}

	return cnt
}