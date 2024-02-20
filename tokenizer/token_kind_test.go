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
