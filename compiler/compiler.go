package compiler

import "fmt"

// This is required by the 'go generate' command
//go:generate goyacc -v y.output -o lexer.go -p brainfuck lexer.y

type Node interface {
	Type() int
}

type CommandList struct {
	children []Node
}

func (n *CommandList) Type() int {
	return 0
}

// increment data pointer <n> values
type IncPtr struct {
	n int
}

func (n *IncPtr) Type() int {
	return 0
}

type BrainfuckParser struct {
	lexer *lexer
	root Node // will hold the parse tree
}

type lexer struct {
	s string
	pos int
}

// This is the main driver / interaction point with yacc
func (l *BrainfuckParser) Lex(lval *brainfuckSymType) int {

	// position in the input
	if l.lexer.pos == len(l.lexer.s) {
		// ... return '0' if we have reached end of input
		return 0
	}

	// grab the byte at current position
	b := l.lexer.s[l.lexer.pos]
	fmt.Printf("l.lexer.s[%d] == %q\n", l.lexer.pos, b)

	switch b {
	case '.':
		// Set the byte value
		lval.b = int('.')
		// move the lexer to the next position in the input string
		l.lexer.pos++
		// return the associated token
		return WRITE_BYTE
	case ',':
		lval.b = int(',')
		l.lexer.pos++
		return READ_BYTE
	case '+':
		lval.b = int('+')
		l.lexer.pos++
		return INC_BYTE
	case '-':
		lval.b = int('-')
		l.lexer.pos++
		return DEC_BYTE
	case '>':
		l.lexer.pos++
		return INC_PTR
	case '<':
		l.lexer.pos++
		return DEC_PTR
	default:
		//fmt.Printf("Default b ? %v\n", b)
		l.lexer.pos++;
		if b == ']' || b == '[' {
			// return the default value (literals use the ASCII code point)
			return int(b)
		} else {
			return OTHER
		}
	}
	//return -1;
}

// Required as part of the yyLexer interface
func (l *BrainfuckParser) Error(s string) {
	fmt.Printf("error: %s\n", s)
}

func Compile(prog string) {
	fmt.Printf("yacc test run for '%s'\n", prog)

	// extra verbosity and debugging
	brainfuckDebug = 1
	brainfuckErrorVerbose = true

	i := brainfuckParse(&BrainfuckParser{
		lexer:lexer{s: prog},
	})

	// non-zero exit code indicates failure
	fmt.Printf("Result ? %d", i)

	if i == 0 {
		fmt.Println("success")
	}
}