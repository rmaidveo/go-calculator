package tokenizer

import (
	"errors"
	"fmt"
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
	stateCtx := newStateContext()
	for index, character := range text {
		switch {
		case unicode.IsSpace(character):
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			continue
		case unicode.IsDigit(character) || character == decimalPointCharacter:
			if character == decimalPointCharacter {
				token, err := stateCtx.createIdentifierToken(index)
				if err != nil && !errors.Is(err, errNoToken) {
					return nil, fmt.Errorf("unable to create a identifier token: %w", err)
				}
				if err == nil {
					tokens = append(tokens, token)
				}
			}

			if err := stateCtx.addCharacterToNumber(index, character); err != nil {
				return nil, fmt.Errorf("unable to add a character to the number: %w", err)
			}
		case unicode.IsLetter(character) || character == '_':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			stateCtx.addCharacterToIdentifier(character)
		case character == '+':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     PlusToken,
				Position: index,
			})
		case character == '-':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     MinusToken,
				Position: index,
			})
		case character == '*':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     AsteriskToken,
				Position: index,
			})
		case character == '/':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     SlashToken,
				Position: index,
			})
		case character == '%':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     PercentToken,
				Position: index,
			})
		case character == '^':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     ExponentiationToken,
				Position: index,
			})
		case character == '(':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     LeftParenthesisToken,
				Position: index,
			})
		case character == ')':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     RightParenthesisToken,
				Position: index,
			})
		case character == ',':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			tokens = append(tokens, Token{
				Kind:     CommaToken,
				Position: index,
			})
		default:
			return nil, fmt.Errorf("unknown character %q at position %d", character, index)
		}
	}

	endOfTextIndex := len(text)
	token, err := stateCtx.createNumberToken(endOfTextIndex)
	if err != nil && !errors.Is(err, errNoToken) {
		return nil, fmt.Errorf("unable to create a number token: %w", err)
	}
	if err == nil {
		tokens = append(tokens, token)
	}

	token, err = stateCtx.createIdentifierToken(endOfTextIndex)
	if err != nil && !errors.Is(err, errNoToken) {
		return nil, fmt.Errorf("unable to create a identifier token: %w", err)
	}
	if err == nil {
		tokens = append(tokens, token)
	}

	return tokens, nil
}
