// lexer.go
package lexer

import (
	"fmt"
	"strings"
)

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

func (lexer *Lexer) Advance() {
	lexer.pos++

	if lexer.pos >= len(lexer.sourceCode) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.sourceCode[lexer.pos]
	}
}

func (lexer *Lexer) Eat() byte {
	c := lexer.currentChar
	lexer.Advance()
	return c
}

func (lexer *Lexer) Peek() byte {
	np := lexer.pos + 1
	if np >= len(lexer.sourceCode) {
		return 0
	}
	return lexer.sourceCode[np]
}

func (lexer *Lexer) Tokenize() []Token {
	var tokens []Token

	add := func(tType TokenType, val string) {
		tokens = append(tokens, Token{Type: tType, Value: val})
	}

	lexer.Advance()

	for lexer.currentChar != 0 {
		switch lexer.currentChar {

		case '+':
			add(PlusToken, string(lexer.Eat()))
			continue
		case '-':
			add(MinusToken, string(lexer.Eat()))
			continue
		case '*':
			add(StarToken, string(lexer.Eat()))
			continue
		case '%':
			add(PercentToken, string(lexer.Eat()))
			continue
		case '(':
			add(LParenToken, string(lexer.Eat()))
			continue
		case ')':
			add(RParenToken, string(lexer.Eat()))
			continue
		case '{':
			add(LBraceToken, string(lexer.Eat()))
			continue
		case '}':
			add(RBraceToken, string(lexer.Eat()))
			continue
		case '[':
			add(LBracketToken, string(lexer.Eat()))
			continue
		case ']':
			add(RBracketToken, string(lexer.Eat()))
			continue
		case ';':
			add(SemicolonToken, string(lexer.Eat()))
			continue
		case ',':
			add(CommaToken, string(lexer.Eat()))
			continue
		case '.':
			add(DotToken, string(lexer.Eat()))
			continue
		case '~':
			add(BitwiseNotToken, string(lexer.Eat()))
			continue
		case '/':
			if lexer.Peek() == '/' {
				lexer.Eat()
				lexer.Eat()
				for lexer.currentChar != 0 && lexer.currentChar != '\n' {
					lexer.Advance()
				}
				continue
			}
			add(SlashToken, string(lexer.Eat()))
			continue

		case '=':
			lexer.Advance()
			if lexer.currentChar == '=' {
				lexer.Advance()
				add(DoubleEqualToken, "==")
			} else {
				add(EqualToken, "=")
			}
			continue
		case '<':
			lexer.Advance()
			if lexer.currentChar == '=' {
				lexer.Advance()
				add(LessEqualToken, "<=")
			} else {
				add(LessThanToken, "<")
			}
			continue
		case '>':
			lexer.Advance()
			if lexer.currentChar == '=' {
				lexer.Advance()
				add(GreaterEqualToken, ">=")
			} else {
				add(GreaterThanToken, ">")
			}
			continue
		case '!':
			lexer.Advance()
			if lexer.currentChar == '=' {
				lexer.Advance()
				add(NotEqualToken, "!=")
			} else {
				add(LogicalNotToken, "!")
			}
			continue
		case '|':
			lexer.Advance()
			if lexer.currentChar == '|' {
				lexer.Advance()
				add(LogicalOrToken, "||")
			} else {
				add(BitwiseOrToken, "|")
			}
			continue
		case '&':
			lexer.Advance()
			if lexer.currentChar == '&' {
				lexer.Advance()
				add(LogicalAndToken, "&&")
			} else {
				add(BitwiseAndToken, "&")
			}
			continue

		case ' ', '\t', '\n', '\r':
			lexer.Advance()
			continue

		case '"':
			lexer.Eat()
			var b strings.Builder
			
			for lexer.currentChar != 0 && lexer.currentChar != '"' {
				if lexer.currentChar == '\\' && lexer.Peek() != 0 {
					lexer.Eat()
					switch lexer.currentChar {
					case 'n':
						b.WriteByte('\n')
					case 't':
						b.WriteByte('\t')
					case '"':
						b.WriteByte('"')
					case '\\':
						b.WriteByte('\\')
					default:
						b.WriteByte(lexer.currentChar)
					}
					lexer.Advance()
					continue
				}
				b.WriteByte(lexer.currentChar)
				lexer.Advance()
			}
			if lexer.currentChar == '"' {
				lexer.Eat()
			}
			add(StringToken, b.String())
			continue

		default:
			if isDigit(lexer.currentChar) {
				var b strings.Builder
				
				for isDigit(lexer.currentChar) {
					b.WriteByte(lexer.currentChar)
					lexer.Advance()
				}

				if lexer.currentChar == '.' && isDigit(lexer.Peek()) {
					b.WriteByte('.')
					lexer.Advance()

					for isDigit(lexer.currentChar) {
						b.WriteByte(lexer.currentChar)
						lexer.Advance()
					}
				}

				add(NumberToken, b.String())
				continue
			}

			if isLetter(lexer.currentChar) {
				var b strings.Builder
				for isLetter(lexer.currentChar) || isDigit(lexer.currentChar) {
					b.WriteByte(lexer.currentChar)
					lexer.Advance()
				}
				idStr := b.String()
				switch idStr {
				case "var":
					add(VarToken, idStr)
				case "if":
					add(IfToken, idStr)
				case "else":
					add(ElseToken, idStr)
				case "for":
					add(ForToken, idStr)
				case "while":
					add(WhileToken, idStr)
				case "func":
					add(FuncToken, idStr)
				case "return":
					add(ReturnToken, idStr)
				default:
					add(IdentifierToken, idStr)
				}
				continue
			}

			fmt.Printf("unrecognized character '%c' at pos %d\n", lexer.currentChar, lexer.pos)
			lexer.Advance()
		}
	}

	add(EOFToken, "EOF")
	return tokens
}
