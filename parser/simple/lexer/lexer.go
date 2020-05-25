// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/nfk93/gocap/parser/simple/token"
)

const (
	NoState    = -1
	NumStates  = 86
	NumSymbols = 110
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: ';'
1: 'c'
2: 'a'
3: 'p'
4: 'c'
5: 'h'
6: 'a'
7: 'n'
8: 'c'
9: 'h'
10: 'a'
11: 'n'
12: 'c'
13: 'o'
14: 'n'
15: 's'
16: 't'
17: 'f'
18: 'u'
19: 'n'
20: 'c'
21: 'i'
22: 'm'
23: 'p'
24: 'o'
25: 'r'
26: 't'
27: 'i'
28: 'n'
29: 't'
30: 'e'
31: 'r'
32: 'f'
33: 'a'
34: 'c'
35: 'e'
36: 'i'
37: 'n'
38: 't'
39: 'm'
40: 'a'
41: 'k'
42: 'e'
43: 'p'
44: 'a'
45: 'c'
46: 'k'
47: 'a'
48: 'g'
49: 'e'
50: 's'
51: 'r'
52: 'i'
53: 'n'
54: 'g'
55: 's'
56: 't'
57: 'r'
58: 'u'
59: 'c'
60: 't'
61: 't'
62: 'y'
63: 'p'
64: 'e'
65: 'v'
66: 'a'
67: 'r'
68: '('
69: ')'
70: '.'
71: ','
72: '*'
73: '_'
74: '.'
75: '.'
76: '.'
77: '{'
78: '}'
79: ':'
80: '='
81: '<'
82: '-'
83: '-'
84: '_'
85: '\'
86: 'a'
87: 'b'
88: 'f'
89: 'n'
90: 'r'
91: 't'
92: 'v'
93: '\'
94: '''
95: '`'
96: '`'
97: '`'
98: '"'
99: '"'
100: ' '
101: '\n'
102: '\t'
103: '\r'
104: 'a'-'z'
105: 'A'-'Z'
106: 'a'-'z'
107: 'A'-'Z'
108: '0'-'9'
109: .
*/
