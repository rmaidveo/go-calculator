package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	PlusToken TokenKind = iota
	MinusToken
	AsteriskToken
	SlashToken
	PercentToken
	ExponentiationToken
	NumberToken
	IdentifierToken
	LeftParenthesisToken
	RightParenthesisToken
	CommaToken
)

type Token struct {
	Kind     TokenKind
	Value    string
	Position int
}

func Tokenize(text string) ([]Token, error) {
	var tokens []Token
	for index, character := range text {
		switch {
		case unicode.IsSpace(character):
			continue
		case character == '+':
			tokens = append(tokens, Token{
				Kind:     PlusToken,
				Position: index,
			})
		case character == '-':
			tokens = append(tokens, Token{
				Kind:     MinusToken,
				Position: index,
			})
		case character == '*':
			tokens = append(tokens, Token{
				Kind:     AsteriskToken,
				Position: index,
			})
		case character == '/':
			tokens = append(tokens, Token{
				Kind:     SlashToken,
				Position: index,
			})
		case character == '%':
			tokens = append(tokens, Token{
				Kind:     PercentToken,
				Position: index,
			})
		case character == '^':
			tokens = append(tokens, Token{
				Kind:     ExponentiationToken,
				Position: index,
			})
		case character == '(':
			tokens = append(tokens, Token{
				Kind:     LeftParenthesisToken,
				Position: index,
			})
		case character == ')':
			tokens = append(tokens, Token{
				Kind:     RightParenthesisToken,
				Position: index,
			})
		case character == ',':
			tokens = append(tokens, Token{
				Kind:     CommaToken,
				Position: index,
			})
		default:
			return nil, fmt.Errorf("unknown character %q at position %d", character, index)
		}
	}

	return tokens, nil
}
