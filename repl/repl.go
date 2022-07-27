package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dev001hajipro/json_parser/lexer"
	"github.com/dev001hajipro/json_parser/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	
	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned { 
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break;
			}
		}
	}
}