package brainfuck

import (

	"testing"
	"fmt"
	"time"
)


func TestInputTokens(t *testing.T) {
	expected := TokenList{Inc, Dec, IncPointer, DecPointer, Write, Read, JumpForward, JumpBackward}

	// The Stringer implementation should output the ASCII code
	fmt.Println(expected)

	ch := Tokenize(expected.String())

	for _, e := range expected {
		token := <- ch
		if Token(e) != LexedToken(token).token {
			fmt.Printf("Expected %v, Got %v(%T)\n", e, token, token)
			t.FailNow()
		}
	}
}

func TestIgnored(t *testing.T) {
	// These should all be tokenized as ignored
	input := "foobar 87shdf87shdf asd"
	ch := Tokenize(input)
	seen := 0

	time := time.After(1 * time.Second)
	for {

		select {
		case token := <-ch:
			seen++
			if LexedToken(token).token != Ignored {
				t.FailNow()
			}
		case <-time:
			expected := len(input)
			if seen != expected {
				fmt.Printf("exepcted %d, observed %d\n", expected, seen)
				t.FailNow()
			} else {
				fmt.Println("Ignored OK!")
				return
			}
		}
	}
}

func TestHelloWorld(t *testing.T) {
	ch := Tokenize("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")

	for t := range ch {
		fmt.Println(t)
	}

	t.FailNow()
}