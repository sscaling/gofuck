//line lexer.y:4
package compiler

import __yyfmt__ "fmt"

//line lexer.y:4
//line lexer.y:8
type brainfuckSymType struct {
	yys int
	b   int // byte
}

const INC_PTR = 57346
const DEC_PTR = 57347
const OTHER = 57348
const INC_BYTE = 57349
const DEC_BYTE = 57350
const READ_BYTE = 57351
const WRITE_BYTE = 57352

var brainfuckToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INC_PTR",
	"DEC_PTR",
	"OTHER",
	"INC_BYTE",
	"DEC_BYTE",
	"READ_BYTE",
	"WRITE_BYTE",
	"'['",
	"']'",
}
var brainfuckStatenames = [...]string{}

const brainfuckEofCode = 1
const brainfuckErrCode = 2
const brainfuckInitialStackSize = 16

//line lexer.y:43

/* Program */

//line yacctab:1
var brainfuckExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const brainfuckNprod = 14
const brainfuckPrivate = 57344

var brainfuckTokenNames []string
var brainfuckStates []string

const brainfuckLast = 43

var brainfuckAct = [...]int{

	9, 10, 11, 5, 6, 7, 8, 4, 16, 9,
	10, 11, 5, 6, 7, 8, 4, 15, 9, 10,
	11, 5, 6, 7, 8, 4, 3, 2, 13, 12,
	1, 0, 0, 0, 0, 14, 0, 0, 0, 0,
	0, 13, 12,
}
var brainfuckPact = [...]int{

	14, 14, -1000, -1000, 5, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -4, -1000, -1000,
}
var brainfuckPgo = [...]int{

	0, 30, 27, 26,
}
var brainfuckR1 = [...]int{

	0, 1, 1, 1, 1, 2, 2, 3, 3, 3,
	3, 3, 3, 3,
}
var brainfuckR2 = [...]int{

	0, 1, 2, 1, 2, 3, 2, 1, 1, 1,
	1, 1, 1, 1,
}
var brainfuckChk = [...]int{

	-1000, -1, -2, -3, 11, 7, 8, 9, 10, 4,
	5, 6, -2, -3, -1, 12, 12,
}
var brainfuckDef = [...]int{

	0, -2, 1, 3, 0, 7, 8, 9, 10, 11,
	12, 13, 2, 4, 0, 6, 5,
}
var brainfuckTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 11, 3, 12,
}
var brainfuckTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var brainfuckTok3 = [...]int{
	0,
}

var brainfuckErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	brainfuckDebug        = 0
	brainfuckErrorVerbose = false
)

type brainfuckLexer interface {
	Lex(lval *brainfuckSymType) int
	Error(s string)
}

type brainfuckParser interface {
	Parse(brainfuckLexer) int
	Lookahead() int
}

type brainfuckParserImpl struct {
	lval  brainfuckSymType
	stack [brainfuckInitialStackSize]brainfuckSymType
	char  int
}

func (p *brainfuckParserImpl) Lookahead() int {
	return p.char
}

func brainfuckNewParser() brainfuckParser {
	return &brainfuckParserImpl{}
}

const brainfuckFlag = -1000

func brainfuckTokname(c int) string {
	if c >= 1 && c-1 < len(brainfuckToknames) {
		if brainfuckToknames[c-1] != "" {
			return brainfuckToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func brainfuckStatname(s int) string {
	if s >= 0 && s < len(brainfuckStatenames) {
		if brainfuckStatenames[s] != "" {
			return brainfuckStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func brainfuckErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !brainfuckErrorVerbose {
		return "syntax error"
	}

	for _, e := range brainfuckErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + brainfuckTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := brainfuckPact[state]
	for tok := TOKSTART; tok-1 < len(brainfuckToknames); tok++ {
		if n := base + tok; n >= 0 && n < brainfuckLast && brainfuckChk[brainfuckAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if brainfuckDef[state] == -2 {
		i := 0
		for brainfuckExca[i] != -1 || brainfuckExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; brainfuckExca[i] >= 0; i += 2 {
			tok := brainfuckExca[i]
			if tok < TOKSTART || brainfuckExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if brainfuckExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += brainfuckTokname(tok)
	}
	return res
}

func brainfucklex1(lex brainfuckLexer, lval *brainfuckSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = brainfuckTok1[0]
		goto out
	}
	if char < len(brainfuckTok1) {
		token = brainfuckTok1[char]
		goto out
	}
	if char >= brainfuckPrivate {
		if char < brainfuckPrivate+len(brainfuckTok2) {
			token = brainfuckTok2[char-brainfuckPrivate]
			goto out
		}
	}
	for i := 0; i < len(brainfuckTok3); i += 2 {
		token = brainfuckTok3[i+0]
		if token == char {
			token = brainfuckTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = brainfuckTok2[1] /* unknown char */
	}
	if brainfuckDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", brainfuckTokname(token), uint(char))
	}
	return char, token
}

func brainfuckParse(brainfucklex brainfuckLexer) int {
	return brainfuckNewParser().Parse(brainfucklex)
}

func (brainfuckrcvr *brainfuckParserImpl) Parse(brainfucklex brainfuckLexer) int {
	var brainfuckn int
	var brainfuckVAL brainfuckSymType
	var brainfuckDollar []brainfuckSymType
	_ = brainfuckDollar // silence set and not used
	brainfuckS := brainfuckrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	brainfuckstate := 0
	brainfuckrcvr.char = -1
	brainfucktoken := -1 // brainfuckrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		brainfuckstate = -1
		brainfuckrcvr.char = -1
		brainfucktoken = -1
	}()
	brainfuckp := -1
	goto brainfuckstack

ret0:
	return 0

ret1:
	return 1

brainfuckstack:
	/* put a state and value onto the stack */
	if brainfuckDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", brainfuckTokname(brainfucktoken), brainfuckStatname(brainfuckstate))
	}

	brainfuckp++
	if brainfuckp >= len(brainfuckS) {
		nyys := make([]brainfuckSymType, len(brainfuckS)*2)
		copy(nyys, brainfuckS)
		brainfuckS = nyys
	}
	brainfuckS[brainfuckp] = brainfuckVAL
	brainfuckS[brainfuckp].yys = brainfuckstate

brainfucknewstate:
	brainfuckn = brainfuckPact[brainfuckstate]
	if brainfuckn <= brainfuckFlag {
		goto brainfuckdefault /* simple state */
	}
	if brainfuckrcvr.char < 0 {
		brainfuckrcvr.char, brainfucktoken = brainfucklex1(brainfucklex, &brainfuckrcvr.lval)
	}
	brainfuckn += brainfucktoken
	if brainfuckn < 0 || brainfuckn >= brainfuckLast {
		goto brainfuckdefault
	}
	brainfuckn = brainfuckAct[brainfuckn]
	if brainfuckChk[brainfuckn] == brainfucktoken { /* valid shift */
		brainfuckrcvr.char = -1
		brainfucktoken = -1
		brainfuckVAL = brainfuckrcvr.lval
		brainfuckstate = brainfuckn
		if Errflag > 0 {
			Errflag--
		}
		goto brainfuckstack
	}

brainfuckdefault:
	/* default state action */
	brainfuckn = brainfuckDef[brainfuckstate]
	if brainfuckn == -2 {
		if brainfuckrcvr.char < 0 {
			brainfuckrcvr.char, brainfucktoken = brainfucklex1(brainfucklex, &brainfuckrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if brainfuckExca[xi+0] == -1 && brainfuckExca[xi+1] == brainfuckstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			brainfuckn = brainfuckExca[xi+0]
			if brainfuckn < 0 || brainfuckn == brainfucktoken {
				break
			}
		}
		brainfuckn = brainfuckExca[xi+1]
		if brainfuckn < 0 {
			goto ret0
		}
	}
	if brainfuckn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			brainfucklex.Error(brainfuckErrorMessage(brainfuckstate, brainfucktoken))
			Nerrs++
			if brainfuckDebug >= 1 {
				__yyfmt__.Printf("%s", brainfuckStatname(brainfuckstate))
				__yyfmt__.Printf(" saw %s\n", brainfuckTokname(brainfucktoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for brainfuckp >= 0 {
				brainfuckn = brainfuckPact[brainfuckS[brainfuckp].yys] + brainfuckErrCode
				if brainfuckn >= 0 && brainfuckn < brainfuckLast {
					brainfuckstate = brainfuckAct[brainfuckn] /* simulate a shift of "error" */
					if brainfuckChk[brainfuckstate] == brainfuckErrCode {
						goto brainfuckstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if brainfuckDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", brainfuckS[brainfuckp].yys)
				}
				brainfuckp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if brainfuckDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", brainfuckTokname(brainfucktoken))
			}
			if brainfucktoken == brainfuckEofCode {
				goto ret1
			}
			brainfuckrcvr.char = -1
			brainfucktoken = -1
			goto brainfucknewstate /* try again in the same state */
		}
	}

	/* reduction by production brainfuckn */
	if brainfuckDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", brainfuckn, brainfuckStatname(brainfuckstate))
	}

	brainfucknt := brainfuckn
	brainfuckpt := brainfuckp
	_ = brainfuckpt // guard against "declared and not used"

	brainfuckp -= brainfuckR2[brainfuckn]
	// brainfuckp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if brainfuckp+1 >= len(brainfuckS) {
		nyys := make([]brainfuckSymType, len(brainfuckS)*2)
		copy(nyys, brainfuckS)
		brainfuckS = nyys
	}
	brainfuckVAL = brainfuckS[brainfuckp+1]

	/* consult goto table to find next state */
	brainfuckn = brainfuckR1[brainfuckn]
	brainfuckg := brainfuckPgo[brainfuckn]
	brainfuckj := brainfuckg + brainfuckS[brainfuckp].yys + 1

	if brainfuckj >= brainfuckLast {
		brainfuckstate = brainfuckAct[brainfuckg]
	} else {
		brainfuckstate = brainfuckAct[brainfuckj]
		if brainfuckChk[brainfuckstate] != -brainfuckn {
			brainfuckstate = brainfuckAct[brainfuckg]
		}
	}
	// dummy call; replaced with literal code
	switch brainfucknt {

	}
	goto brainfuckstack /* stack new state and value */
}
