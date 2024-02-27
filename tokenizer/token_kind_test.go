package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTokenKind(t *testing.T) {
	type args struct {
		character rune
	}

	tests := []struct {
		name    string
		args    args
		want    TokenKind
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success/+",
			args:    args{character: '+'},
			want:    PlusToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/-",
			args:    args{character: '-'},
			want:    MinusToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/*",
			args:    args{character: '*'},
			want:    AsteriskToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success//",
			args:    args{character: '/'},
			want:    SlashToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/%",
			args:    args{character: '%'},
			want:    PercentToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/^",
			args:    args{character: '^'},
			want:    ExponentiationToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/(",
			args:    args{character: '('},
			want:    LeftParenthesisToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/)",
			args:    args{character: ')'},
			want:    RightParenthesisToken,
			wantErr: assert.NoError,
		},
		{
			name:    "success/,",
			args:    args{character: ','},
			want:    CommaToken,
			wantErr: assert.NoError,
		},
		{
			name:    "error/@",
			args:    args{character: '@'},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTokenKind(tt.args.character)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}

func TestTokenKind_Precedence(t *testing.T) {
	tests := []struct {
		name string
		kind TokenKind
		want int
	}{
		{name: "+", kind: PlusToken, want: 1},
		{name: "-", kind: MinusToken, want: 1},
		{name: "*", kind: AsteriskToken, want: 2},
		{name: "/", kind: SlashToken, want: 2},
		{name: "%", kind: PercentToken, want: 2},
		{name: "^", kind: ExponentiationToken, want: 3},
		{name: "number", kind: NumberToken, want: 0},
		{name: "identifier", kind: IdentifierToken, want: 0},
		{name: "(", kind: LeftParenthesisToken, want: 0},
		{name: ")", kind: RightParenthesisToken, want: 0},
		{name: ",", kind: CommaToken, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.kind.Precedence()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTokenKind_IsOperator(t *testing.T) {
	tests := []struct {
		name string
		kind TokenKind
		want assert.BoolAssertionFunc
	}{
		{name: "+", kind: PlusToken, want: assert.True},
		{name: "-", kind: MinusToken, want: assert.True},
		{name: "*", kind: AsteriskToken, want: assert.True},
		{name: "/", kind: SlashToken, want: assert.True},
		{name: "%", kind: PercentToken, want: assert.True},
		{name: "^", kind: ExponentiationToken, want: assert.True},
		{name: "number", kind: NumberToken, want: assert.False},
		{name: "identifier", kind: IdentifierToken, want: assert.False},
		{name: "(", kind: LeftParenthesisToken, want: assert.False},
		{name: ")", kind: RightParenthesisToken, want: assert.False},
		{name: ",", kind: CommaToken, want: assert.False},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.kind.IsOperator()

			tt.want(t, got)
		})
	}
}

func TestTokenKind_String(t *testing.T) {
	tests := []struct {
		name string
		kind TokenKind
		want string
	}{
		{name: "+", kind: PlusToken, want: "+"},
		{name: "-", kind: MinusToken, want: "-"},
		{name: "*", kind: AsteriskToken, want: "*"},
		{name: "/", kind: SlashToken, want: "/"},
		{name: "%", kind: PercentToken, want: "%"},
		{name: "^", kind: ExponentiationToken, want: "^"},
		{name: "number", kind: NumberToken, want: ""},
		{name: "identifier", kind: IdentifierToken, want: ""},
		{name: "(", kind: LeftParenthesisToken, want: "("},
		{name: ")", kind: RightParenthesisToken, want: ")"},
		{name: ",", kind: CommaToken, want: ","},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.kind.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
