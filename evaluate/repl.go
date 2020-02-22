import (
	"bufio"
	"fmt"
	"io"
	"../grammar"
	"../lexer"
)

const PROMPT = "<meow>^..^<meow>"
//helper function for  input and output values 
func Start(in io.Reader, out io.Writer){
	scanner := bufion.Scanner(in)

	for{
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned{
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
	
		for tok := lexer.NextToken();tok.Type != token.EOF; tok = lexer.NextToken(){
			fmt.Printf("%+v\n",tok)
		}
	}


}
