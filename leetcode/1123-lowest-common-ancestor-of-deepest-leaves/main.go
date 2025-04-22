package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode) (int, *TreeNode)

	dfs = func(root *TreeNode) (int, *TreeNode) {
		if root == nil {
			return 0, nil
		}
		leftDeep, leftNode := dfs(root.Left)
		rightDeep, rightNode := dfs(root.Right)

		if leftDeep > rightDeep {
			return leftDeep + 1, leftNode
		}

		if rightDeep > leftDeep {
			return rightDeep + 1, rightNode
		}

		return leftDeep + 1, root
	}

	_, lca := dfs(root)

	return lca
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	// Construct example tree:
	root := newNode(3)
	root.Left = newNode(5)
	root.Right = newNode(1)
	root.Left.Left = newNode(6)
	root.Left.Right = newNode(2)
	root.Right.Right = newNode(8)
	root.Left.Right.Left = newNode(7)
	root.Left.Right.Right = newNode(4)

	lca := lcaDeepestLeaves(root)
	fmt.Println("LCA of deepest leaves:", lca.Val) // Output: 2
}
