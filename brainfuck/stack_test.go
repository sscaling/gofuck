package brainfuck

import (
	"fmt"
	"testing"
)

func TestStack_String(t *testing.T) {
	s := NewStack()
	s.Push(&StackNode{3})
	s.Push(&StackNode{2})
	s.Push(&StackNode{1})

	if "[0]3:[1]2:[2]1:" != fmt.Sprintf("%s", s) {
		t.Fail()
	}
}
