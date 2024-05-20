package main

import (
	"fmt"
	"log"
)

type Node struct {
	data     int
	nextNode *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) push(data int) {
	n := new(Node)
	n.data = data
	n.nextNode = s.top
	s.top = n
	s.size += 1
}

func (s *Stack) pop() int {
	if s.is_empty() {
		log.Fatal("The Stack is empty.")
	}
	popped_node := s.top
	s.top = s.top.nextNode
	s.size -= 1
	return popped_node.data
}

func (s *Stack) is_empty() bool {
	return s.top == nil
}

func (s *Stack) peek() int {
	return s.top.data
}

func main() {
	var s Stack
	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s)
	fmt.Println("Top of the stack: ", s.peek())
	fmt.Println("You just popped: ", s.pop())
	fmt.Println("Top of the stack: ", s.peek())
	fmt.Println("Stack Size: ", s.size)
}
