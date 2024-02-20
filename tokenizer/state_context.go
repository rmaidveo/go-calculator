package tokenizer

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errNoToken = errors.New("no token")
)

type stateContext struct {
	state                 State
	numberHasDecimalPoint bool
	buffer                strings.Builder
}

func newStateContext() stateContext {
	return stateContext{
		state:                 DefaultState,
		numberHasDecimalPoint: false,
	}
}

func (stateCtx *stateContext) addCharacterToNumber(index int, character rune) error {
	if stateCtx.state == DefaultState {
		stateCtx.state = NumberState
	}

	if character == decimalPointCharacter && stateCtx.state == NumberState {
		if stateCtx.numberHasDecimalPoint {
			return fmt.Errorf("duplicate decimal point in the number at position %d", index)
		}

		stateCtx.numberHasDecimalPoint = true
	}

	stateCtx.buffer.WriteRune(character)

	return nil
}

func (stateCtx *stateContext) addCharacterToIdentifier(character rune) {
	if stateCtx.state == DefaultState {
		stateCtx.state = IdentifierState
	}

	stateCtx.buffer.WriteRune(character)
}

func (stateCtx *stateContext) createNumberToken(index int) (Token, error) {
	if stateCtx.state != NumberState {
		return Token{}, errNoToken
	}

	value := stateCtx.buffer.String()
	position := index - len(value)
	if value == string(decimalPointCharacter) {
		return Token{}, fmt.Errorf("the number has only a decimal point at position %d", position)
	}

	stateCtx.state = DefaultState
	stateCtx.numberHasDecimalPoint = false
	stateCtx.buffer.Reset()

	token := Token{
		Kind:     NumberToken,
		Value:    value,
		Position: position,
	}
	return token, nil
}

func (stateCtx *stateContext) createIdentifierToken(index int) (Token, error) {
	if stateCtx.state != IdentifierState {
		return Token{}, errNoToken
	}

	value := stateCtx.buffer.String()

	stateCtx.state = DefaultState
	stateCtx.buffer.Reset()

	token := Token{
		Kind:     IdentifierToken,
		Value:    value,
		Position: index - len(value),
	}
	return token, nil
}
