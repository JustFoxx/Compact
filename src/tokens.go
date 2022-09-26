package main

type TokenType string

const (
	STRING           TokenType = "STRING"
	NUMBER           TokenType = "NUMBER"
	NULL             TokenType = "NULL"
	VAR              TokenType = "VAR"
	PRINT            TokenType = "PRINT"
	SCAN             TokenType = "SCAN"
	ASSIGN           TokenType = "ASSIGN"
	EQUAL            TokenType = "EQUAL"
	NOT_EQUAL        TokenType = "NOT_EQUAL"
	MINUS            TokenType = "MINUS"
	PLUS             TokenType = "PLUS"
	BRACE_OPEN       TokenType = "BRACE_OPEN"
	BRACE_CLOSE      TokenType = "BRACE_CLOSE"
	PAREN_OPEN       TokenType = "PAREN_OPEN"
	PAREN_CLOSE      TokenType = "PAREN_CLOSE"
	MOD              TokenType = "MOD"
	MULTIPLE         TokenType = "MULTIPLE"
	DIVIDE           TokenType = "DIVIDE"
	THIS             TokenType = "THIS"
	SEMICOLON        TokenType = "SEMICOLON"
	IF               TokenType = "IF"
	ELSE             TokenType = "ELSE"
	FALSE            TokenType = "FALSE"
	COMMENT          TokenType = "COMMENT"
	AND              TokenType = "AND"
	OR               TokenType = "OR"
	BIGGER_OR_EQUAL  TokenType = "BIGGER_OR_EQUAL"
	SMALLER_OR_EQUAL TokenType = "SMALLER_OR_EQUAL"
	BIGGER           TokenType = "BIGGER"
	SMALLER          TokenType = "SMALLER"
	STRUCT_OPEN      TokenType = "STRUCT_OPEN"
	STRUCT_CLOSE     TokenType = "STRUCT_CLOSE"
	ARRAY_OPEN       TokenType = "ARRAY_OPEN"
	ARRAY_CLOSE      TokenType = "ARRAY_CLOSE"
	LIST_OPEN        TokenType = "LIST_OPEN"
	LIST_CLOSE       TokenType = "LIST_CLOSE"
	COMMA            TokenType = "COMMA"
	INCREASE         TokenType = "INCREASE"
	DECREASE         TokenType = "DECREASE"
	POWER            TokenType = "POWER"
	ROOT             TokenType = "ROOT"
	ABSOLUTE         TokenType = "ABSOLUTE"
	ROUND            TokenType = "ROUND"
	RETURN           TokenType = "RETURN"
	DOT              TokenType = "DOT"
	EOL              TokenType = "EOL"
	IDENTIFIER       TokenType = "IDENTIFIER"
)