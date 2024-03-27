package calculator

import (
	"testing"

	"github.com/rmaidveo/go-calculator/evaluator"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	type args struct {
		text      string
		variables map[string]float64
		functions map[string]evaluator.Function
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				text:      "23 + 42",
				variables: map[string]float64{},
				functions: map[string]evaluator.Function{
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
			name: "error/unable to tokenize",
			args: args{
				text:      "23 @ 42",
				variables: map[string]float64{},
				functions: map[string]evaluator.Function{
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
			name: "error/unable to translate",
			args: args{
				text:      "(23 + 42",
				variables: map[string]float64{},
				functions: map[string]evaluator.Function{
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
			name: "error/unable to evaluate",
			args: args{
				text:      "23 +",
				variables: map[string]float64{},
				functions: map[string]evaluator.Function{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.text, tt.args.variables, tt.args.functions)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
