/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	st, ed := root, root
	for st != nil && st.Left != nil {
		var last *Node
		for n := st; n != nil; n = n.Next {
			if last != nil {
				last.Next = n.Left
			}
			last = n.Right

			n.Left.Next = n.Right
		}
		st, ed = st.Left, ed.Right
	}
	return root
}