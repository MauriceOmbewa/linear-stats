package maths

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		sum    float64
		length float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Average with positive integers",
			args: args{
				sum:    15,
				length: 3,
			},
			want: 5,
		},

		{
			name: "Test Average with negative integers",
			args: args{
				sum:    -15,
				length: 3,
			},
			want: -5,
		},

		{
			name: "Test Average with one element",
			args: args{
				sum:    15,
				length: 1,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.sum, tt.args.length); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}
