package brainfuck

import "fmt"

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func (n *Node) Equals(i int) bool {
	return n != nil && n.Value == i
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []*Node
	count int
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Peek() *Node {
	if s.count == 0 {
		return nil
	}
	return s.nodes[s.count - 1]
}

func (s *Stack) Size() int {
	return s.count
}
