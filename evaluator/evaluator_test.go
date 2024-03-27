package evaluator

import (
	"testing"
	"testing/iotest"

	"github.com/rmaidveo/go-calculator/translator"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		commands  []translator.Command
		variables map[string]float64
		functions map[string]Function
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success/push number",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "23", Position: 42},
				},
				variables: map[string]float64{},
				functions: map[string]Function{},
			},
			want:    23,
			wantErr: assert.NoError,
		},
		{
			name: "success/push variable",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushVariableCommand, Operand: "x", Position: 42},
				},
				variables: map[string]float64{"x": 23},
				functions: map[string]Function{},
			},
			want:    23,
			wantErr: assert.NoError,
		},
		{
			name: "success/call function",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "23", Position: 123},
					{Kind: translator.PushNumberCommand, Operand: "42", Position: 142},
					{Kind: translator.CallFunctionCommand, Operand: "+", Position: 150},
				},
				variables: map[string]float64{},
				functions: map[string]Function{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			want:    65,
			wantErr: assert.NoError,
		},
		{
			name: "error/push number",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "invalid-number", Position: 42},
				},
				variables: map[string]float64{},
				functions: map[string]Function{},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error/push variable",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushVariableCommand, Operand: "unknown-variable", Position: 42},
				},
				variables: map[string]float64{},
				functions: map[string]Function{},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error/call function/unknown function",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "23", Position: 123},
					{Kind: translator.PushNumberCommand, Operand: "42", Position: 142},
					{Kind: translator.CallFunctionCommand, Operand: "unknown-function", Position: 150},
				},
				variables: map[string]float64{},
				functions: map[string]Function{},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error/call function/number stack is empty for an argument",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "23", Position: 123},
					{Kind: translator.CallFunctionCommand, Operand: "+", Position: 150},
				},
				variables: map[string]float64{},
				functions: map[string]Function{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error/call function/unable to call the function",
			args: args{
				commands: []translator.Command{
					{Kind: translator.PushNumberCommand, Operand: "23", Position: 123},
					{Kind: translator.PushNumberCommand, Operand: "42", Position: 142},
					{Kind: translator.CallFunctionCommand, Operand: "+", Position: 150},
				},
				variables: map[string]float64{},
				functions: map[string]Function{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return 0, iotest.ErrTimeout
						},
					},
				},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error/no commands",
			args: args{
				commands:  []translator.Command{},
				variables: map[string]float64{},
				functions: map[string]Function{},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.args.commands, tt.args.variables, tt.args.functions)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
