package main

import (
	"fmt"
)

type BST struct {
	val   int
	left  *BST
	right *BST
}

type OperationsBst interface {
	Insert(root *BST, val int) *BST
	//Remove() *BST
	PrintTree(root *BST)
}

func (bst BST) Insert(root *BST, val int) *BST {

	temp := root
	newNode := &BST{val: val, left: nil, right: nil}

	if temp == nil {
		return newNode
	}

	for temp != nil {

		if val > temp.val && temp.right != nil {
			temp = temp.right
		} else if val > temp.val && temp.right == nil {
			temp.right = newNode
			break
		} else if val < temp.val && temp.left != nil {
			temp = temp.left
		} else if val < temp.val && temp.left == nil {
			temp.left = newNode
			break
		}
	}

	return root

}

func (bst BST) PrintTree(node *BST) {

	if node == nil {
		return
	}
	bst.PrintTree(node.left)
	fmt.Print(node.val)
	bst.PrintTree(node.right)

}

func (bst BST) DeleteNode(root *BST, key int) *BST {

	if root == nil {
		fmt.Println("Empty structure")
		return root
	}

	if key < root.val {

		bst.DeleteNode(root.left, key)

	} else if key > root.val {

		bst.DeleteNode(root.right, key)

	} else {

		//check if node exists in the left and node exists in the right

		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			root.val = minvalue(root.right)
			root.right = bst.DeleteNode(root.right, root.val)
		}

	}

	return root

}

//minimum value of the tree
func minvalue(node *BST) int {

	if node.left != nil {
		minvalue(node.left)
	}
	return node.val

}
