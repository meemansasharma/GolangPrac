package main

import "fmt"

/**
A structure to define a node
functions :
insert
delete
print
reverse
*/
//Single link list

type Node struct {
	next *Node
	val  int
}

//insert function
func Insert(head *Node, val int) *Node {

	//create new node
	newNode := &Node{}
	newNode.next = nil
	newNode.val = val
	if head == nil {
		head = newNode
		return head
	}
	//traverse the end position
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	//add new node at the end
	temp.next = newNode
	return head
}

func Print(head *Node) {

	//check if the node is empty
	if head == nil {
		fmt.Println("Empty list")
	} else {
		//traverse and print the list
		temp := head
		for temp != nil {
			fmt.Println(temp.val)
			temp = temp.next
		}
	}
}

//delete function
func Delete(head *Node, val int) *Node {

	if head == nil {
		return head
	}
	if head.val == val {
		return head.next
	}
	//traverse the end position
	temp := head
	prev := head
	for temp != nil {
		//find the value
		if temp.val == val && temp.next != nil {
			prev.next = temp.next
			break
		} else if temp.val == val {
			prev.next = nil
			break
		}
		prev = temp
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("Not Found")
	}
	return head
}

//Reverse Link List
func ReverseList(head *Node) *Node {

	if head == nil {
		fmt.Println("Empty List")
		return head
	}
	current := head
	var prev *Node
	for current != nil {

		nextNode := current.next
		current.next = prev
		prev = current
		current = nextNode

	}

	head = prev
	return head
}
