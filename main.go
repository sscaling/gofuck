package main

import (
	"fmt"
	bf "github.com/sscaling/gofuck/brainfuck"
)
//a
const a = ">  ++++++++++[>++++++++++><<-]>---><<>."
// ... is this the equivalent of ?
const a2 = ">++++++++++[>++++++++++<-]>---."
//12
const simple = "++[>+<-]>."
const helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func main() {
	fmt.Println("Start")

	bf.Compile(helloWorld)

	fmt.Println("Finish")
}
