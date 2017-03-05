package brainfuck

// Contains input character position and BrainFuck token
type LexedToken struct {
	position int
	token Token
}


// Takes a ASCII bytes and writes LexedTokens sequentially to channel
func Tokenize(program string) <-chan LexedToken {

	ch := make(chan LexedToken)

	go func() {
		defer close(ch)

		for i := 0; i < len(program); i++ {
			ch <- LexedToken{i, GetToken(program[i])}
		}
	}()

	return ch
}