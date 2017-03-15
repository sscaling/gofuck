package brainfuck

import (
	"bytes"
	"fmt"
)

type Token int
type TokenList []Token

const (
	Ignored      = Token(0) // anything that doesn't match the characters below
	IncPointer   = Token(1)
	DecPointer   = Token(2)
	Inc          = Token(3)
	Dec          = Token(4)
	Write        = Token(5)
	Read         = Token(6)
	JumpForward  = Token(7)
	JumpBackward = Token(8)
)

var tokens = map[byte]Token{
	'>': IncPointer,
	'<': DecPointer,
	'+': Inc,
	'-': Dec,
	'.': Write,
	',': Read,
	'[': JumpForward,
	']': JumpBackward,
}

var chars = map[Token]byte{
	IncPointer:   '>',
	DecPointer:   '<',
	Inc:          '+',
	Dec:          '-',
	Write:        '.',
	Read:         ',',
	JumpForward:  '[',
	JumpBackward: ']',
}

func GetToken(b byte) Token {
	if t, found := tokens[b]; found {
		return t
	}

	return Ignored
}

func (tokens TokenList) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for _, t := range tokens {
		if t != Ignored {
			buffer.WriteByte(chars[t])
		}
	}
	return buffer.String()
}

func (t Token) String() string {
	if c, found := chars[t]; found {
		return fmt.Sprintf("%q", c)
	}

	return "?"
}
