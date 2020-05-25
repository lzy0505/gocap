// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/nfk93/gocap/parser/simple/token"
)

const (
	NoState    = -1
	NumStates  = 89
	NumSymbols = 117
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
41: 'p'
42: 'm'
43: 'a'
44: 'k'
45: 'e'
46: 'p'
47: 'a'
48: 'c'
49: 'k'
50: 'a'
51: 'g'
52: 'e'
53: 's'
54: 'r'
55: 'i'
56: 'n'
57: 'g'
58: 's'
59: 't'
60: 'r'
61: 'u'
62: 'c'
63: 't'
64: 't'
65: 'y'
66: 'p'
67: 'e'
68: 'v'
69: 'a'
70: 'r'
71: '('
72: ')'
73: '['
74: ']'
75: '{'
76: '}'
77: '.'
78: ','
79: '*'
80: '<'
81: '-'
82: '<'
83: '-'
84: '-'
85: ':'
86: '='
87: '_'
88: '.'
89: '.'
90: '.'
91: '_'
92: '\'
93: 'a'
94: 'b'
95: 'f'
96: 'n'
97: 'r'
98: 't'
99: 'v'
100: '\'
101: '''
102: '`'
103: '`'
104: '`'
105: '"'
106: '"'
107: ' '
108: '\n'
109: '\t'
110: '\r'
111: 'a'-'z'
112: 'A'-'Z'
113: 'a'-'z'
114: 'A'-'Z'
115: '0'-'9'
116: .
*/
