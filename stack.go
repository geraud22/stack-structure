package main

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

func main() {
	var s Stack
	s.push(10)
	s.push(20)
	s.push(30)
}
