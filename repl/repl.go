package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yuzuy/yoru/evaluator"
	"github.com/yuzuy/yoru/lexer"
	"github.com/yuzuy/yoru/object"
	"github.com/yuzuy/yoru/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const monkeyFace = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []error) {
	io.WriteString(out, monkeyFace)
	io.WriteString(out, "Woops! we run into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, err := range errors {
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
