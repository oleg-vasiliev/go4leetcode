package lc45

import (
	"testing"
)

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	type testCase struct {
		name string
		args args
		want int
	}

	tests := []testCase{
		{
			name: "Ex1",
			args: args{
				[]int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "Ex2",
			args: args{
				[]int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
