/* Simple Singly Linked List */
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node = new(Node)

func add(d int) {
	var new_node *Node = new(Node)
	new_node.data = d

	if head == nil {
		head = new_node
		head.next = nil
	} else {
		new_node.next = head
		head = new_node
	}
}

func print() {
	var current *Node = new(Node)
	current = head

	if current == nil {
		fmt.Println("List is empty")
	} else {
		fmt.Println("List of Elements :")
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func main() {
	head = nil
	add(1)
	add(3)
	add(2)
	print()
}
