package brainfuck

import (
	"fmt"
	"io"
)

const EXECUTION_LIMIT = 20000

func Interpret(program string, in io.ByteReader, out io.ByteWriter) {
	// lex
	tokens := Tokenize(program)

	// data for program
	bytes := make([]byte, 30000)

	// the pointer
	p := 0

	instructions := make([]LexedToken, len(program))
	instructionCount := 0

	// FIXME: horrible, unnecessary use of channels leads to this ugliness
	for t := range tokens {
		lexed := LexedToken(t)

		if lexed.token != Ignored {
			instructions[instructionCount] = lexed
			instructionCount++
		}
	}

	fmt.Printf("%v (%d/%d)\n", instructions, len(instructions), cap(instructions))

	// FIXME: also horrible. More playing with some Go basics. Should remove.
	stack := NewStack()
	executionCtr := 0
	instructionPtr := 0

	for instructionPtr < instructionCount {

		executionCtr++
		if executionCtr > EXECUTION_LIMIT {
			fmt.Printf("ERROR: Exceeded Execution limit %d. Infinite loop?\n", EXECUTION_LIMIT)
			return
		}

		// Should we move to the next contiguous instruction?
		increment := true
		instruction := instructions[instructionPtr]

		//fmt.Printf("P=(%d) ?, I=(%d) %v, S={%v}, D=%v\n", p, instructionPtr, instruction.token, stack, bytes[0:10])

		switch instruction.token {
		case Inc:
			bytes[p]++
		case Dec:
			bytes[p]--
		case IncPointer:
			p++
		case DecPointer:
			p--
		case Read:
			if b, err := in.ReadByte(); err == nil {
				bytes[p] = b
			} else {
				fmt.Printf("ERROR: reading input (%v)\n", err)
				return
			}

		case Write:
			//fmt.Printf("Write '%v'\n", bytes[p])
			if err := out.WriteByte(bytes[p]); err != nil {
				fmt.Printf("ERROR: writing output (%v)\n", err)
				return
			}

		case JumpForward:
			//fmt.Println("JumpForward")
			if bytes[p] == 0 {
				found := false
				depth := 0
				for !found {

					instructionPtr++

					if instructionPtr < 0 || instructionPtr > instructionCount {
						fmt.Printf("ERROR: Invalid instructionPtr %d\n", instructionPtr)
						return
					}

					instruction = instructions[instructionPtr]

					//fmt.Printf("i:%d, d:%d, t:%v\n", instructionPtr, depth, instruction)

					switch instruction.token {
					case JumpForward:
						depth++
						//fmt.Printf("JumpFoward ++depth %d\n", depth)
					case JumpBackward:
						if 0 != depth {
							depth--
							//fmt.Printf("JumpBackwards --depth %d\n", depth)
						} else {
							//fmt.Println("Break")
							found = true
							break
						}
					}

				}
			} else {
				//fmt.Printf("Fwd Push %d\n", instructionPtr)
				// push this instruction location onto the stack
				stack.Push(&StackNode{instructionPtr})
			}
		case JumpBackward:
			//fmt.Printf("JumpBackward %d\n", p)

			n := stack.Pop()

			if bytes[p] == 0 {
				// break out of loop, and continue
				//fmt.Printf("Pop Break ? %v\n", n)
			} else {

				//fmt.Printf("Pop Rewind to ? %v (%T)\n", n, n)
				// go back to previous instruction

				if n == nil {
					fmt.Printf("ERROR: Nothing on stack to pop")
					return
				}

				increment = false
				instructionPtr = n.Value
			}
		default:
			fmt.Printf("%v (%T) not supported\n", instruction.token, instruction.token)
		}

		if increment {
			instructionPtr++
		}
	}
}
