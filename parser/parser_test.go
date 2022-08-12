package parser

import (
	"testing"

	"github.com/dev001hajipro/json_parser/lexer"
)

func TestParse(t *testing.T) {
	input := `["hello", -123, 10.5, ["world", 5] ]`

	l := lexer.New(input)
	p := New(l)
	json := p.Parse()
	ls, ok := json.([]any)
	if !ok {
		t.Fatalf("json not []any. got=%T", json)
	}
	if len(ls) != 4 {
		t.Fatalf("ls is not currect size. got=%d", len(ls))
	}

	tests := []any{"hello", -123, 10.5}
	checkElementCast(ls, tests, t)


	lls1, ok := ls[3].([]any)
	if !ok {
		t.Fatalf("json not []any. got=%T", lls1)
	}
	tests = []any{"world", 5}
	checkElementCast(lls1, tests, t)

}

func checkElementCast(ls []any, tests []any, t *testing.T) {
	for i, expectedValue := range tests {
		switch v := expectedValue.(type) {
		case string:
			if ls[i].(string) != v {
				t.Fatalf("(string)value is not '%v'. got=%v", v, ls[i])
			}
		case bool:
			if ls[i].(bool) != v {
				t.Fatalf("(string)value is not '%v'. got=%v", v, ls[i])
			}
		case int:
			if int(ls[i].(int64)) != v {
				t.Fatalf("(int)value is not '%v'. got=%v", v, ls[i])
			}
		case float64:
			if ls[i] != v {
				t.Fatalf("(float64)value is not '%v'. got=%v", v, ls[i])
			}
		default:
			t.Fatalf("cast error.")
		}
	}
}

func TestParseObject(t *testing.T) {
	input := `{"hello":"world"}`

	l := lexer.New(input)
	p := New(l)
	json := p.Parse()
	ls, ok := json.(map[interface{}]interface{}) // map[any]any
	if !ok {
		t.Fatalf("json not map[any]any. got=%T", json)
	}

	if len(ls) != 1 {
		t.Fatalf("json does not containn 1 object. got=%d", len(ls))
	}

	if ls["hello"] != "world" {
		t.Fatalf("json does not containn 'hello' key. got=%v", ls["hello"])
	}

}
