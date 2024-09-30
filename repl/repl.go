package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Hien-Trinh/interpreter-go/evaluator"
	"github.com/Hien-Trinh/interpreter-go/lexer"
	"github.com/Hien-Trinh/interpreter-go/object"
	"github.com/Hien-Trinh/interpreter-go/parser"
)

const PROMPT = ">> "
const MONKEY_FACE = `
           __,__
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |'  | |
| \   \  \ 0 | 0 /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |   \._ _./   |
      \    \'~'/    /
       '._ '-=-' _.'
          '-----'
`

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
			printParseError(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			if err != nil {
				return
			}

			_, err = io.WriteString(out, "\n")
			if err != nil {
				return
			}
		}
	}
}

func printParseError(out io.Writer, errors []string) {
	_, err := io.WriteString(out, MONKEY_FACE)
	if err != nil {
		return
	}

	_, err = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	if err != nil {
		return
	}

	_, err = io.WriteString(out, " parser errors:\n")
	if err != nil {
		return
	}

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
