class Solution {

    fun dfs(now: TreeNode?, v: Int, current: Int, target: Int) {
        val n = now ?: return
        if (current == target) {
            n.left = TreeNode(v).apply { left = n.left }
            n.right = TreeNode(v).apply { right = n.right }
            return
        }

        dfs(n.left, v, current + 1, target)
        dfs(n.right, v, current + 1, target)
    }

    fun addOneRow(root: TreeNode?, `val`: Int, depth: Int): TreeNode? {
        if (depth == 1) return TreeNode(`val`).apply { left = root }
        dfs(root, `val`, 2, depth)
        return root
    }
}