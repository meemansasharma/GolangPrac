package main

import "fmt"

func (bst BST) InsertDL(head *BST, val int) *BST {

	temp := head
	newNode := &BST{val: val, left: nil, right: nil}

	if temp == nil {
		return newNode
	}

	for temp.right != nil {
		temp = temp.right
	}

	temp.right = newNode
	newNode.left = temp
	return head
}

func (bst BST) Remove(head *BST, val int) *BST {

	temp := head

	if temp == nil {
		return nil
	} else if temp.val == val {
		temp = temp.right
		temp.left = nil
		return temp
	}

	for temp != nil {
		nodeL := temp.left
		nodeR := temp.right
		if temp.val == val && nodeL != nil && nodeR != nil {
			nodeL.right = temp.right
			nodeR.left = nodeL
			break
		} else if temp.val == val && nodeR == nil {
			nodeL.right = nil
		}
		temp = temp.right
	}

	return head
}
func (bst BST) ShowList(node *BST) {

	if node == nil {
		return
	}
	for node != nil {
		fmt.Print(node.val, " ")
		node = node.right
	}

}

func (bst BST) ReverseDL(head *BST) *BST {

	var temp *BST
	current := head

	if current == nil || current.right == nil {
		return current
	}
	for current != nil {
		temp = current.left
		current.left = current.right
		current.right = temp
		current = current.left
	}

	if temp != nil {
		head = temp.left
	}
	return head
}
