package parser

import (
	"testing"

	"github.com/dev001hajipro/json_parser/lexer"
)

func TestParse(t *testing.T) {
	
	l := lexer.New(`["hello", -123, 10.5, ["world", 5] ]`)
	p := New(l)
	json := p.Parse()
	ls := json.([]any)
	t.Logf("type=%T. value=%v\n", ls[0], ls[0])
	t.Logf("type=%T. value=%v\n", ls[1], ls[1])
	t.Logf("type=%T. value=%v\n", ls[2], ls[2])
	t.Logf("type=%T. value=%v\n", ls[3], ls[3])
	lls1 := ls[3].([]any)
	t.Logf("type=%T. value=%v\n", lls1[0], lls1[0])
	t.Logf("type=%T. value=%v\n", lls1[1], lls1[1])
}

func TestParseObject(t *testing.T) {
	
	l := lexer.New(`{"hello":"world}`)
	p := New(l)
	json := p.Parse()
	ls := json.(map[any]any)
	t.Logf("type=%T. \n", ls)
}
