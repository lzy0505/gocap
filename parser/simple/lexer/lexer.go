// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/nfk93/gocap/parser/simple/token"
)

const (
	NoState    = -1
	NumStates  = 206
	NumSymbols = 296
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
0: '\n'
1: ';'
2: 'b'
3: 'r'
4: 'e'
5: 'a'
6: 'k'
7: 'c'
8: 'a'
9: 'p'
10: 'c'
11: 'h'
12: 'a'
13: 'n'
14: 'c'
15: 'a'
16: 's'
17: 'e'
18: 'c'
19: 'h'
20: 'a'
21: 'n'
22: 'c'
23: 'o'
24: 'n'
25: 's'
26: 't'
27: 'c'
28: 'o'
29: 'n'
30: 't'
31: 'i'
32: 'n'
33: 'u'
34: 'e'
35: 'd'
36: 'e'
37: 'f'
38: 'a'
39: 'u'
40: 'l'
41: 't'
42: 'd'
43: 'e'
44: 'f'
45: 'e'
46: 'r'
47: 'e'
48: 'l'
49: 's'
50: 'e'
51: 'f'
52: 'a'
53: 'l'
54: 'l'
55: 't'
56: 'h'
57: 'r'
58: 'o'
59: 'u'
60: 'g'
61: 'h'
62: 'f'
63: 'o'
64: 'r'
65: 'f'
66: 'u'
67: 'n'
68: 'c'
69: 'g'
70: 'o'
71: 'g'
72: 'o'
73: 't'
74: 'o'
75: 'i'
76: 'f'
77: 'i'
78: 'm'
79: 'p'
80: 'o'
81: 'r'
82: 't'
83: 'i'
84: 'n'
85: 't'
86: 'e'
87: 'r'
88: 'f'
89: 'a'
90: 'c'
91: 'e'
92: 'i'
93: 'n'
94: 't'
95: 'm'
96: 'a'
97: 'k'
98: 'e'
99: 'm'
100: 'a'
101: 'p'
102: 'p'
103: 'a'
104: 'c'
105: 'k'
106: 'a'
107: 'g'
108: 'e'
109: 'r'
110: 'a'
111: 'n'
112: 'g'
113: 'e'
114: 'r'
115: 'e'
116: 't'
117: 'u'
118: 'r'
119: 'n'
120: 's'
121: 'e'
122: 'l'
123: 'e'
124: 'c'
125: 't'
126: 's'
127: 't'
128: 'r'
129: 'u'
130: 'c'
131: 't'
132: 's'
133: 't'
134: 'r'
135: 'i'
136: 'n'
137: 'g'
138: 's'
139: 'w'
140: 'i'
141: 't'
142: 'c'
143: 'h'
144: 't'
145: 'y'
146: 'p'
147: 'e'
148: 'v'
149: 'a'
150: 'r'
151: '.'
152: '.'
153: '.'
154: '('
155: ')'
156: '['
157: ']'
158: '{'
159: '}'
160: '.'
161: ','
162: ':'
163: '+'
164: '-'
165: '*'
166: '/'
167: '%'
168: '&'
169: '|'
170: '^'
171: '<'
172: '<'
173: '>'
174: '>'
175: '&'
176: '^'
177: '+'
178: '='
179: '-'
180: '='
181: '*'
182: '='
183: '/'
184: '='
185: '%'
186: '='
187: '&'
188: '='
189: '|'
190: '='
191: '^'
192: '='
193: '<'
194: '<'
195: '='
196: '>'
197: '>'
198: '='
199: '&'
200: '^'
201: '='
202: '&'
203: '&'
204: '|'
205: '|'
206: '<'
207: '-'
208: '<'
209: '-'
210: '<'
211: '-'
212: '-'
213: '<'
214: '+'
215: '-'
216: '+'
217: '+'
218: '-'
219: '-'
220: '='
221: '='
222: '<'
223: '>'
224: '='
225: '!'
226: '!'
227: '='
228: '<'
229: '='
230: '>'
231: '='
232: ':'
233: '='
234: '_'
235: '.'
236: '.'
237: '.'
238: '0'
239: '1'
240: '_'
241: '_'
242: '_'
243: '_'
244: '0'
245: '_'
246: '0'
247: 'b'
248: 'B'
249: '_'
250: '0'
251: 'o'
252: 'O'
253: '_'
254: '0'
255: 'x'
256: 'X'
257: '_'
258: '_'
259: '\'
260: 'a'
261: 'b'
262: 'f'
263: 'n'
264: 'r'
265: 't'
266: 'v'
267: '\'
268: '''
269: '`'
270: '`'
271: '`'
272: '"'
273: '"'
274: '/'
275: '/'
276: '\n'
277: '/'
278: '*'
279: '*'
280: '/'
281: ' '
282: '\t'
283: '\r'
284: 'a'-'z'
285: 'A'-'Z'
286: '0'-'9'
287: '0'-'7'
288: '0'-'9'
289: 'a'-'f'
290: 'A'-'F'
291: '1'-'9'
292: 'a'-'z'
293: 'A'-'Z'
294: '0'-'9'
295: .
*/
