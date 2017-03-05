package brainfuck

import (
	"bytes"
)

type Token int
type TokenList []Token

const Ignored = Token(0) // anything that doesn't match the characters below
const IncPointer = Token(1)
const DecPointer = Token(2)
const Inc = Token(3)
const Dec = Token(4)
const Write = Token(5)
const Read = Token(6)
const JumpForward = Token(7)
const JumpBackward = Token(8)

var tokens = map[byte]Token {
	'>': IncPointer,
	'<': DecPointer,
	'+': Inc,
	'-': Dec,
	'.': Write,
	',': Read,
	'[': JumpForward,
	']': JumpBackward,
}

var chars = map[Token]byte {
	IncPointer: '>',
	DecPointer: '<',
	Inc: '+',
	Dec: '-',
	Write: '.',
	Read: ',',
	JumpForward: '[',
	JumpBackward: ']',
}

func GetToken(b byte) Token {
	if t, found := tokens[b]; found {
		return t
	}

	return Ignored
}

func (tokens TokenList) String() string {
	buffer := bytes.NewBuffer([]byte {})
	for _, t := range tokens {
		if t != Ignored {
			buffer.WriteByte(chars[t])
		}
	}
	return buffer.String()
}