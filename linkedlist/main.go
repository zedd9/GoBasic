package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func makeNode(v int) *Node {
	return &Node{val: v}
}

func (n *Node) setNext(next *Node) {
	n.next = n
}

type LinkedList struct {
	header *Node
	tail   *Node
}

func (list *LinkedList) PrintAll() {
	if list.header == nil {
		return
	}

	var node *Node = list.header
	for ; node.next != nil; node = node.next {
		fmt.Print(node.val, " -> ")
	}

	fmt.Println(node.val)
}

func (list *LinkedList) addItem(val int) {
	newNode := makeNode(val)

	if list.header == nil {
		list.header = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList) removeItem(val int) {
	if list.header == nil {
		return
	}

	if list.header.val == val {
		if list.header == list.tail {
			list.header = nil
			list.tail = nil
		} else {
			list.header = list.header.next
		}

		return
	}

	var prevNode *Node = list.header
	var curNode *Node = prevNode.next
	for ; curNode != nil; curNode = curNode.next {
		if curNode.val == val {
			if curNode == list.tail {
				prevNode.next = nil
				list.tail = prevNode
			} else {
				prevNode.next = curNode.next
			}

			return
		}

		prevNode = curNode
	}
}

func main() {
	list := LinkedList{header: nil, tail: nil}

	list.addItem(1)
	list.addItem(2)
	list.addItem(3)
	list.addItem(4)

	list.PrintAll()

	list.removeItem(1)
	//list.removeItem(2)
	list.removeItem(3)
	//list.removeItem(4)
	list.addItem(3)
	list.addItem(4)

	list.PrintAll()
}
