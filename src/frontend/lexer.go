package frontend

import (
	"fmt"
	"strings"
)

// Lexer holds the source and current position state.
type Lexer struct {
	sourceCode  string
	pos         int
	currentChar byte
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func NewLexer(sourceCode string) *Lexer {
	return &Lexer{
		sourceCode: sourceCode,
		pos:        -1,
	}
}

// Advance moves the cursor forward and updates currentChar (0 == EOF).
func (l *Lexer) Advance() {
	l.pos++
	if l.pos >= len(l.sourceCode) {
		l.currentChar = 0
	} else {
		l.currentChar = l.sourceCode[l.pos]
	}
}

// Eat returns the current char then advances once.
func (l *Lexer) Eat() byte {
	c := l.currentChar
	l.Advance()
	return c
}

// Peek returns the next char without advancing (or 0 at EOF).
func (l *Lexer) Peek() byte {
	np := l.pos + 1
	if np >= len(l.sourceCode) {
		return 0
	}
	return l.sourceCode[np]
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token

	// small helper to append tokens cleanly
	add := func(tType TokenType, val string) {
		tokens = append(tokens, Token{Type: tType, Value: val})
	}

	l.Advance() // initialize currentChar

	for l.currentChar != 0 {
		switch l.currentChar {

		// single-char operators and punctuation
		case '+':
			add(Plus, string(l.Eat()))
			continue
		case '-':
			add(Minus, string(l.Eat()))
			continue
		case '*':
			add(Star, string(l.Eat()))
			continue
		case '%':
			add(Percent, string(l.Eat()))
			continue
		case '(':
			add(LParen, string(l.Eat()))
			continue
		case ')':
			add(RParen, string(l.Eat()))
			continue
		case '{':
			add(LBrace, string(l.Eat()))
			continue
		case '}':
			add(RBrace, string(l.Eat()))
			continue
		case '[':
			add(LBracket, string(l.Eat()))
			continue
		case ']':
			add(RBracket, string(l.Eat()))
			continue
		case ';':
			add(Semicolon, string(l.Eat()))
			continue
		case ',':
			add(Comma, string(l.Eat()))
			continue
		case '.':
			add(Dot, string(l.Eat()))
			continue

		// slash could be division or line comment //
		case '/':
			if l.Peek() == '/' {
				// consume the '//' and skip to end of line
				l.Eat() // '/'
				l.Eat() // second '/'
				for l.currentChar != 0 && l.currentChar != '\n' {
					l.Advance()
				}
				continue
			}
			add(Slash, string(l.Eat()))
			continue

		// comparisons and multi-char operators
		case '=':
			l.Advance()
			if l.currentChar == '=' {
				l.Advance()
				add(DoubleEqual, "==")
			} else {
				add(Equal, "=")
			}
			continue
		case '<':
			l.Advance()
			if l.currentChar == '=' {
				l.Advance()
				add(LessEqual, "<=")
			} else {
				add(LessThan, "<")
			}
			continue
		case '>':
			l.Advance()
			if l.currentChar == '=' {
				l.Advance()
				add(GreaterEqual, ">=")
			} else {
				add(GreaterThan, ">")
			}
			continue
		case '!':
			l.Advance()
			if l.currentChar == '=' {
				l.Advance()
				add(NotEqual, "!=")
			} else {
				add(Not, "!")
			}
			continue

		// whitespace -> skip
		case ' ', '\t', '\n', '\r':
			l.Advance()
			continue

		// string literal (double quotes) with basic escapes
		case '"':
			l.Eat() // consume opening quote
			var b strings.Builder
			for l.currentChar != 0 && l.currentChar != '"' {
				if l.currentChar == '\\' && l.Peek() != 0 {
					l.Eat() // consume backslash
					switch l.currentChar {
					case 'n':
						b.WriteByte('\n')
					case 't':
						b.WriteByte('\t')
					case '"':
						b.WriteByte('"')
					case '\\':
						b.WriteByte('\\')
					default:
						// unknown escape: keep literal char
						b.WriteByte(l.currentChar)
					}
					l.Advance() // move after escaped char
					continue
				}
				b.WriteByte(l.currentChar)
				l.Advance()
			}
			if l.currentChar == '"' {
				l.Eat() // consume closing quote
			}
			add(String, b.String())
			continue

		default:
			// numbers (integers + optional fractional part)
			if isDigit(l.currentChar) {
				var b strings.Builder
				for isDigit(l.currentChar) {
					b.WriteByte(l.currentChar)
					l.Advance()
				}
				if l.currentChar == '.' && isDigit(l.Peek()) {
					b.WriteByte('.')
					l.Advance()
					for isDigit(l.currentChar) {
						b.WriteByte(l.currentChar)
						l.Advance()
					}
				}
				add(Number, b.String())
				continue
			}

			// identifiers & keywords
			if isLetter(l.currentChar) {
				var b strings.Builder
				for isLetter(l.currentChar) || isDigit(l.currentChar) {
					b.WriteByte(l.currentChar)
					l.Advance()
				}
				idStr := b.String()
				switch idStr {
				case "var":
					add(Var, idStr)
				case "const":
					add(Const, idStr)
				case "if":
					add(If, idStr)
				case "else":
					add(Else, idStr)
				case "for":
					add(For, idStr)
				case "while":
					add(While, idStr)
				case "func":
					add(Func, idStr)
				case "return":
					add(Return, idStr)
				default:
					add(Identifier, idStr)
				}
				continue
			}

			// unrecognized -> warn and advance so we don't loop forever
			fmt.Printf("Unrecognized character '%c' at pos %d\n", l.currentChar, l.pos)
			l.Advance()
		}
	}

	add(EOF, "EOF")
	return tokens
}
