package main

import "fmt"

func Reverse(newString string) string {

	runes := []rune(newString)
	rLength := len(newString) - 1
	lLength := 0
	for lLength < rLength {
		runes[lLength], runes[rLength] = runes[rLength], runes[lLength]
		lLength++
		rLength--
	}
	return string(runes)
}
func main() {
	// fmt.Println(Reverse("meemansa meaning is to explain"))
	// fmt.Println(MergeSort([]int{2, 4, 8, 67}, []int{5, 7, 54, 90}))
	// fmt.Println(RecurringNumber([]int{2, 1, 1, 2, 4, 8, 67}))
	// fmt.Println("Linked list")
	// head := Insert(nil, 2)
	// Print(head)
	// Insert(head, 9)
	// Insert(head, 10)
	// Insert(head, 4)
	// fmt.Println("Linked list after delete")
	// head = Delete(head, 9)
	// Print(head)
	// fmt.Println("Linked list Reverse")
	// head = ReverseList(head)
	// Print(head)

	//--Binary tree --
	fmt.Println("Binary tree Operations")
	bst := BST{}
	root := bst.Insert(nil, 10)
	bst.Insert(root, 3)
	bst.Insert(root, -1)
	bst.Insert(root, 6)
	bst.Insert(root, 90)
	bst.Insert(root, 70)
	bst.PrintTree(root)
	//reference -- https://www.geeksforgeeks.org/binary-search-tree-set-2-delete/
	fmt.Println("\n Binary tree Operations-- delete")
	bst.DeleteNode(root, 3)
	bst.PrintTree(root)

	// //levelorderTraversal
	// levelOrderTraversal()
	//Double link list
	// fmt.Println(" Double link list Operations")
	// list := BST{}
	// ListHead := list.InsertDL(nil, 2)
	// list.InsertDL(ListHead, 3)
	// list.InsertDL(ListHead, -1)
	// list.ShowList(ListHead)
	// fmt.Println("Delete Operations")
	// list.ShowList(list.Remove(ListHead, -1))
	// list.InsertDL(ListHead, 7)
	// fmt.Println("show reversed Operations")
	// list.ShowList(list.ReverseDL(ListHead))
	// graph := Graph{}
	// graph.Creategraph()
	// graph.getVertexCount()
}
