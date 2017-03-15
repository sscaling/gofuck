package brainfuck

import "fmt"

type StackNode struct {
	Value int
}

func (n *StackNode) String() string {
	return fmt.Sprint(n.Value)
}

func (n *StackNode) Equals(i int) bool {
	return n != nil && n.Value == i
}

type Stack struct {
	nodes []*StackNode
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) String() string {
	str := ""
	for i := 0; i < s.count; i++ {
		str = fmt.Sprintf("%s[%d]%v:", str, i, s.nodes[i])
	}
	return str
}

func (s *Stack) Push(n *StackNode) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *StackNode {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Peek() *StackNode {
	if s.count == 0 {
		return nil
	}
	return s.nodes[s.count-1]
}

func (s *Stack) Size() int {
	return s.count
}
