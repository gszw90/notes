package test

import (
	"github.com/gszw90/notes/sort_algorithm"
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{nums: []int{1, 1, 5, 2, 3}},
			want: []int{1, 1, 2, 3, 5},
		},
		{
			name: "test2",
			args: args{nums: []int{1}},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test4",
			args: args{nums: []int{4, 3, 2, 1}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test5",
			args: args{nums: []int{4, 3, 2, 1, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort_algorithm.BubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
