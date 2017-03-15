package main

import (
	"fmt"
	bf "github.com/sscaling/gofuck/brainfuck"
	"os"
)

//input
const input = ",>,>+++[<<++>+>-]<<.>."

//a
const a = ">  ++++++++++[>++++++++++><<-]>---><<>."

// ... is this the equivalent of ?
const a2 = ">++++++++++[>++++++++++<-]>---."

//12
const simple = "++[>+<-]>."
const helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
const rot13 = "-,+[-[>>++++[>++++++++<-]<+<-[>+>+>-[>>>]<[[>+<-]>>+>]<<<<<-]]>>>[-]+>--[-[<->+++[-]]]<[++++++++++++<[>-[>+>>]>[+[<+>-]>+>>]<<<<<-]>>[<+>-]>[-[-<<[-]>>]<<[<<->>-]>>]<<[<<+>>-]]<[-]<.[-]<-,+]"

func main() {
	fmt.Println("Start")

	bf.Interpret(rot13, stdinReader(1), stdoutWriter(1))

	fmt.Println("Finish")
}

// FIXME: what is the idiomatic/better way to do this? use a struct?
type stdinReader int
type stdoutWriter int

func (r stdinReader) ReadByte() (byte, error) {
	buff := make([]byte, 1)

	if _, err := os.Stdin.Read(buff); err == nil {
		return buff[0], nil
	} else {
		return ' ', err
	}
}

func (r stdoutWriter) WriteByte(c byte) error {
	fmt.Printf("%s", string(c))
	return nil
}
