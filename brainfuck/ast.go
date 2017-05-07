package brainfuck

type Node interface {
	Token() Token
}

/**
* essentially should create an AST, something like:
*
*                main
*
*     6     *p>0     p++     out
*             |
*    p++ inc 11 p-- dec (while probably represented as a condition and children)
*
* Which represents a program, that loads 6 into position p[0], while p > 0
* increment p[1] by 11, then set p, back to p[0]. Control is broken when
* the value at p[0] is 0, otherwise the loop continues. 'out' writes to stdout.
* The equivalent brainfuck would be '++++++[>+++++++++++<-]>.'

How would this look as byte code?

//data 30000  // LOAD / STORE work against this
//P 0 // 'p' -> P index register - PIDX_LOAD
PIDX_STORE work against this

L0:
  LOAD // data[p] -  stack[0]
  PUSH 6 // stack[0,6]
  ADD // Add 0 + 6, and pop - stack [6]
  STORE // store at position 0 -> pop 6, P=0, stack[], data[6]

L1:
  PUSH 0  // P=0, stack[0], data[6]
  CMPLE L3 // P <=0 ? goto L3 : L2. pop. P=0, stack[], data[6]

L2: // L(abel)1
  PUSH 1 // stack[1]
  PIDX_LOAD // stack[1,0] NOTE: this is more akin to an Index register: https://en.wikipedia.org/wiki/Index_register
  ADD    // stack[1] pop both, add save result
  PIDX_STORE // P = 1, pop, stack []
  LOAD   // stack[0]
  PUSH 11 // stack[0, 11]
  ADD    // stack [11]
  STORE  // at position 'P', p=1 stack[], data[6,11]
  PIDX_LOAD // stack[1]
  PUSH 1 // stack[1,1]
  SUB    // stack[0]
  PIDX_STORE // P=0, stack[]
  LOAD    // stack[6]
  PUSH 1  // stack[6,1]
  SUB     // stack[5]
  STORE   // P=0, stack[5], data[5,11]
  JMP L1

L3:
  PUSH 1
  PIDX_LOA
  ADD
  PIDX_STORE
  PUTC  // put character at P
*/
