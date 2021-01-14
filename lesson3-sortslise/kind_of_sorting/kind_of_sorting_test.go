package kind_of_sorting

import (
	"reflect"
	"testing"
)

var result []int

func BenchmarkBubbleSort(b *testing.B) {
var res []int
	abc := []int{484, 935, 639, 140, 228}
for i := 0; i < b.N; i++ {
	res = BubbleSort(abc)
}
	result = res
}

func BenchmarkInsertSort(b *testing.B) {
	var res []int
	abc := []int{484, 935, 639, 140, 228}
	for i := 0; i < b.N; i++ {
		res = InsertSort(abc)
	}
	result = res
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Test 1", input: []int{484, 935, 639, 140, 228}, want: []int{140, 228, 484, 639, 935}},
		{name: "Test 2", input: []int{4, 3, 1, 5, 2}, want: []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range tests {
		got := BubbleSort(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestInsertSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Test 1", input: []int{484, 935, 639, 140, 228}, want: []int{140, 228, 484, 639, 935}},
		{name: "Test 2", input: []int{4, 3, 1, 5, 2}, want: []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range tests {
		got := InsertSort(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}