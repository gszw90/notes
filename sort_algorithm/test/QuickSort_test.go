package test

import (
	"github.com/gszw90/notes/sort_algorithm"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{nums: []int{1, 1, 5, 2, 3}, low: 0, high: 4},
			want: []int{1, 1, 2, 3, 5},
		},
		{
			name: "test2",
			args: args{nums: []int{1}, low: 0, high: 0},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{nums: []int{1, 2, 3, 4}, low: 0, high: 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test4",
			args: args{nums: []int{4, 3, 2, 1}, low: 0, high: 3},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test5",
			args: args{nums: []int{4, 3, 2, 1, 5}, low: 0, high: 4},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort_algorithm.QuickSort(tt.args.nums, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
