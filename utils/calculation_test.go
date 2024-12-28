package utils_test

import (
	"testing"
	"web_calculator/utils"
)

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		want     float64
		wantErr  bool
	}{
		{
			name:    "Addition",
			expr:    "2+3",
			want:    5.0,
			wantErr: false,
		},
		{
			name:    "Subtraction",
			expr:    "5-2",
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "Multiplication",
			expr:    "4*3",
			want:    12.0,
			wantErr: false,
		},
		{
			name:    "Division",
			expr:    "8/4",
			want:    2.0,
			wantErr: false,
		},
		{
			name:    "Complex expression",
			expr:    "3+5*2-8/4",
			want:    11.0,
			wantErr: false,
		},
		{
			name:    "Division by zero",
			expr:    "5/0",
			want:    0.0,
			wantErr: true,
		},
		{
			name:    "Invalid characters",
			expr:    "5+a",
			want:    0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.Calc(tt.expr)

			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("EvaluateExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
