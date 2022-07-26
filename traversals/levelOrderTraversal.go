package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal() {

	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	output := levelOrder(&root)
	fmt.Println(output)

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}

	Queue := []*TreeNode{root}
	var result [][]int
	for len(Queue) > 0 { //1,2

		sz := len(Queue) //1,2

		level := []int{}
		for i := 0; i < sz; i++ { //1,2
			node := Queue[0] //=1,2
			Queue = Queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				Queue = append(Queue, node.Left)
			}
			if node.Right != nil {
				Queue = append(Queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
