func averageOfSubtree(root *TreeNode) int {
	var res int
	var dfs func(node *TreeNode) (sum, cnt int)
	dfs = func(node *TreeNode) (sum int, cnt int) {
		if node == nil {
			return
		}
		sL, cL := dfs(node.Left)
		sR, cR := dfs(node.Right)
		sum, cnt = sL+node.Val+sR, cL+1+cR
		if sum/cnt == node.Val {
			res++
		}
		return
	}
	dfs(root)
	return res
}
