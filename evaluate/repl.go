package repl

import (
	"bufio"
	"compiler/lexer"
	parser "compiler/parser"
	"fmt"
	"io"
)

const PROMPT = "<meow>^..^<meow>"
//helper function for  input and output values 
func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)

	for{
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned{
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)
		parser := parser.New(lexer)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}


}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! WE ARE IN TROUBLES!\n")
	io.WriteString(out, "Stack trace of parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
