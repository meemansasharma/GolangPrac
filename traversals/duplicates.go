package main

import (
	"fmt"
)

func (node Node) CheckDuplicates(head *Node) *Node {

	current := head
	temp := head
	if head == nil {
		return head
	}
	//122246
	for temp != nil {
		if current.val == temp.val {
			temp = temp.next
			current.next = temp
		} else {
			current = temp
			temp = temp.next
		}

	}
	return head
}
func (node Node) CreateList() *Node {
	head := Node{}
	head.val = 1
	head.next = nil
	return &head
}

func (node *Node) Insert(head *Node, val int) {
	if head == nil {
		return
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	newNode := Node{
		val:  val,
		next: nil,
	}
	temp.next = &newNode
}

func (node Node) Print(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
