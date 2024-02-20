package tokenizer

import "fmt"

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

func ParseTokenKind(character rune) (TokenKind, error) {
	switch character {
	case '+':
		return PlusToken, nil
	case '-':
		return MinusToken, nil
	case '*':
		return AsteriskToken, nil
	case '/':
		return SlashToken, nil
	case '%':
		return PercentToken, nil
	case '^':
		return ExponentiationToken, nil
	case '(':
		return LeftParenthesisToken, nil
	case ')':
		return RightParenthesisToken, nil
	case ',':
		return CommaToken, nil
	default:
		return 0, fmt.Errorf("unknown character %q", character)
	}
}
