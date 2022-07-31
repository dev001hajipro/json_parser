package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

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
			fmt.Printf("%v\n", tok)
			if tok.Type == token.EOF  {
				break;
			}
			if tok.Type == token.RBRACE {
				break;
			}
			if tok.Type == token.ILLEGAL {
				log.Fatalf("token error: %v\n", tok)
			}
		}
	}
}