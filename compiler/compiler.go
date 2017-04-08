package compiler

import "fmt"

// This is required by the 'go generate' command
//go:generate goyacc -v y.output -o lexer.go -p brainfuck lexer.y

type BrainfuckLex struct {
	s string
	pos int
}

func (l *BrainfuckLex) Lex(lval *brainfuckSymType) int {

	// position in the input
	if l.pos == len(l.s) {
		// ... return '0' if we have reached end of input
		return 0
	}

	// grab the byte at current position
	b := l.s[l.pos]
	fmt.Printf("l.s[%d] == %q\n", l.pos, b)

	switch b {
	case '.':
		// Set the byte value
		lval.b = int('.')
		// move the lexer to the next position in the input string
		l.pos++
		// return the associated token
		return WRITE_BYTE
	case ',':
		lval.b = int(',')
		l.pos++
		return READ_BYTE
	case '+':
		lval.b = int('+')
		l.pos++
		return INC_BYTE
	case '-':
		lval.b = int('-')
		l.pos++
		return DEC_BYTE
	case '>':
		l.pos++
		return INC_PTR
	case '<':
		l.pos++
		return DEC_PTR
	default:
		//fmt.Printf("Default b ? %v\n", b)
		l.pos++;
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
func (l *BrainfuckLex) Error(s string) {
	fmt.Printf("error: %s\n", s)
}

func Compile(prog string) {
	fmt.Printf("yacc test run for '%s'\n", prog)

	// extra verbosity and debugging
	brainfuckDebug = 1
	brainfuckErrorVerbose = true


	i := brainfuckParse(&BrainfuckLex{s:prog})

	// non-zero exit code indicates failure
	fmt.Printf("Result ? %d", i)
}