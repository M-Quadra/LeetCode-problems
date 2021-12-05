func move(n int, ap, bp, cp *[]int) {
	if n < 2 {
		*cp = append(*cp, (*ap)[len(*ap)-1])
		*ap = (*ap)[:len(*ap)-1]
		return
	}

	move(1, ap, cp, bp)
	move(n-1, ap, bp, cp)
	move(1, bp, ap, cp)
}

func hanota(A, B, C []int) []int {
	move(len(A), &A, &B, &C)
	return C
}