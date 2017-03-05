package brainfuck

import "fmt"

func Compile(program string) {

	// lex
	tokens := Tokenize(program)
	// parse - build AST
	//  - walk - build byte-code
	// execute
	//  - vm - execute byte-code

	// data for program
	bytes := make([]byte, 30000)

	// the pointer
	p := 0

	instructions := make([]LexedToken, len(program))
	valid := 0

	for t := range tokens {
		lexed := LexedToken(t)

		if lexed.token != Ignored {
			instructions[valid] = lexed
			valid++
		}
	}

	fmt.Printf("%v (%d/%d)\n", instructions, len(instructions), cap(instructions))

	count := len(instructions[:valid])

	stack := NewStack()

	for i := 0; i < count; i++ {

		//fmt.Println(bytes[0:10])

		instruction := instructions[i]
		switch instruction.token {
		case Inc:
			bytes[p]++
		case Dec:
			bytes[p]--
		case IncPointer:
			p++
		case DecPointer:
			p--
		case Write:
			// TODO: write to stdout
			fmt.Printf("%q", bytes[p])
		case JumpForward:
			if bytes[p] != 0 {
				n := stack.Peek()
				//fmt.Printf("%v (%T)\n", n, n)
				if !n.Equals(i) {
					// push this instruction location onto the stack
					stack.Push(&Node{i})
				}
			}
		case JumpBackward:
			if bytes[p] == 0 {
				// break out of loop, and continue
				stack.Pop()
			}  else {
				n := stack.Peek()
				//fmt.Printf("%v (%T)\n", n, n)
				// go back to previous instruction
				i = n.Value
			}
		case Read:
			fallthrough
		default:
			fmt.Printf("%v (%T) not supported", instruction.token, instruction.token)
		}
	}

	fmt.Println()

	//fmt.Println(bytes[0:10])
}