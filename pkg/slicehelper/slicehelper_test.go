package slicehelper

import (
	"reflect"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	sliceData := []string{"eve", "test", "has data", "lorem", "ipsum"}
	type args struct {
		s string
		a []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "found string in slice", args: args{s: "has data", a: sliceData}, want: true},
		{name: "not found string in slice", args: args{s: "not data in this slice", a: sliceData}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Flatten(t *testing.T) {
	type args struct {
		lists [][]any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "flatten string", args: args{[][]any{{"a", "b", "c"}, {"d", "e", "f"}}}, want: []any{"a", "b", "c", "d", "e", "f"}},
		{name: "flatten number", args: args{[][]any{{1, 2, 3}, {4, 5, 6}}}, want: []any{1, 2, 3, 4, 5, 6}},
		{name: "flatten mix", args: args{[][]any{{1, "a", 2}, {"b", 3, "c"}}}, want: []any{1, "a", 2, "b", 3, "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
