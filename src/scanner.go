package main

import "fmt"

type Scanner struct {
	chars  []byte
	char   int
	start  int
	line   int
	tokens []Token
}

func (s *Scanner) scanTokens() {
	for !s.isEnd() {
		s.start = s.char
		s.scanToken()
	}
	fmt.Println(s.tokens)
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := string(s.chars[s.start:s.char])
	s.tokens = append(s.tokens, Token{tokenType: tokenType, lexeme: text, literal: literal, line: s.line})
}

func (s *Scanner) addClearToken(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) scanToken() TokenType {
	char := s.advance()
	switch char {
	case '-':
		if s.match('0') {
			return NULL
		} else if s.match('-') {
			return DECREASE
		} else {
			return  MINUS
		}
	case '$':
		return VAR
	case '<':
		if s.match('<') {
			return PRINT
		} else if s.match('=') {
			return BIGGER_OR_EQUAL
		} else {
			return BIGGER
		}
	case '>':
		if s.match('>') {
			return SCAN
		} else if s.match('=') {
			return SMALLER_OR_EQUAL
		} else {
			return SMALLER
		}
	case '=':
		if s.match('=') {
			return EQUAL
		} else {
			return ASSIGN
		}
	case '!':
		if s.match('=') {
			return NOT_EQUAL
		} else {
			return FALSE
		}
	case '{':
		return BRACE_OPEN
	case '}':
		return BRACE_CLOSE
	case '(':
		s.addClearToken(PAREN_OPEN)
	case ')':
		s.addClearToken(PAREN_CLOSE)
	case '%':
		s.addClearToken(MOD)
	case '*':
		if s.match('*') {
			s.addClearToken(POWER)
		} else {
			s.addClearToken(MULTIPLE)
		}
	case '/':
		if s.match('/') {
			s.addClearToken(ROOT)
		} else {
			s.addClearToken(DIVIDE)
		}
	case '@':
		if s.match('@') {
			s.addClearToken(RETURN)
		} else {
			s.addClearToken(THIS)
		}
	case '?':
		s.addClearToken(IF)
	case ':':
		s.addClearToken(ELSE)
	case '#':
		s.addClearToken(COMMENT)
	case '&':
		s.addClearToken(AND)
	case '|':
		if s.match('{') {
			s.addClearToken(STRUCT_OPEN)
		} else if s.match('}') {
			s.addClearToken(STRUCT_CLOSE)
		} else {
			s.addClearToken(OR)
		}
	case '[':
		if s.match('[') {
			s.addClearToken(LIST_OPEN)
		} else {
			s.addClearToken(ARRAY_OPEN)
		}
	case ']':
		if s.match(']') {
			s.addClearToken(LIST_CLOSE)
		} else {
			s.addClearToken(ARRAY_CLOSE)
		}
	case ',':
		s.addClearToken(COMMA)
	case '.':
		s.addClearToken(DOT)
	case '~':
		if s.match('~') {
			s.addClearToken(ROUND)
		} else {
			s.addClearToken(ABSOLUTE)
		}
	case '"':
		s.string()
	case ' ', '\r', '\t','\n':
		break
	case ';':
		s.line++
	default:
		if s.isDigit(char) {
			s.number()
		} else if s.isAlpha(char) {
			s.identifier()
		} else {
			e.tokenError(s.line,char, "Unexpected character.")
		}
	}

}

func (s *Scanner) isEnd() bool {
	return len(s.chars) <= s.char
}

func (s *Scanner) advance() byte {
	s.char++
	return s.chars[s.char-1]
}

func (s *Scanner) peek() byte {
	if s.isEnd() {
		return 0
	}
	return s.chars[s.char]
}

func (s *Scanner) peekNext() byte {
	if s.char+1 >= len(s.chars) {
		return 0
	}
	return s.chars[s.char+1]
}

func (s *Scanner) match(expected byte) bool {
	if s.isEnd() {
		return false
	}
	if s.chars[s.char] != expected {
		return false
	}
	s.char++
	return true
}

//string
func (s *Scanner) string() {
	for s.peek() != '"' && !s.isEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isEnd() {
		e.lineError(s.line, "Unterminated string.")
		return
	}
	s.advance()
	value := string(s.chars[s.start+1 : s.char-1])
	s.addToken(STRING, value)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	value := string(s.chars[s.start:s.char])
	s.addToken(NUMBER, value)
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	s.addClearToken(IDENTIFIER)
}
