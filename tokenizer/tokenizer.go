package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	decimalPointCharacter = '.'
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

type State int

const (
	DefaultState State = iota
	NumberState
	IdentifierState
)

func Tokenize(text string) ([]Token, error) {
	var tokens []Token
	state := DefaultState
	numberHasDecimalPoint := false
	var buffer strings.Builder
	for index, character := range text {
		switch {
		case unicode.IsSpace(character):
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			continue
		case unicode.IsDigit(character) || character == decimalPointCharacter:
			if state == DefaultState {
				state = NumberState
			}

			if character == decimalPointCharacter {
				if numberHasDecimalPoint {
					return nil, fmt.Errorf("duplicate decimal point in the number at position %d", index)
				}

				numberHasDecimalPoint = true
			}

			buffer.WriteRune(character)
		case character == '+':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     PlusToken,
				Position: index,
			})
		case character == '-':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     MinusToken,
				Position: index,
			})
		case character == '*':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     AsteriskToken,
				Position: index,
			})
		case character == '/':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     SlashToken,
				Position: index,
			})
		case character == '%':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     PercentToken,
				Position: index,
			})
		case character == '^':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     ExponentiationToken,
				Position: index,
			})
		case character == '(':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     LeftParenthesisToken,
				Position: index,
			})
		case character == ')':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     RightParenthesisToken,
				Position: index,
			})
		case character == ',':
			if state == NumberState {
				state = DefaultState

				value := buffer.String()
				position := index - len(value)
				if value == string(decimalPointCharacter) {
					return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
				}

				tokens = append(tokens, Token{
					Kind:     NumberToken,
					Value:    value,
					Position: position,
				})

				numberHasDecimalPoint = false
				buffer.Reset()
			}

			tokens = append(tokens, Token{
				Kind:     CommaToken,
				Position: index,
			})
		default:
			return nil, fmt.Errorf("unknown character %q at position %d", character, index)
		}
	}

	if state == NumberState {
		value := buffer.String()
		position := len(text) - len(value)
		if value == string(decimalPointCharacter) {
			return nil, fmt.Errorf("the number has only a decimal point at position %d", position)
		}

		tokens = append(tokens, Token{
			Kind:     NumberToken,
			Value:    value,
			Position: position,
		})
	}

	return tokens, nil
}
