package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c/", sep: ",", want: []string{"a/b/c/"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("%v: expected: %#v, got: %#v", name, tc.want, got)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := map[string]struct {
		a    int
		b    int
		want int
	}{
		"testcase1": {a: 1, b: 2, want: 3},
		"testcase2": {a: 11, b: 22, want: 33},
		"testcase3": {a: 10, b: 20, want: 30},
		"testcase4": {a: 1, b: -2, want: -3},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sum(tc.a, tc.b)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("%v: expected: %#v, got: %#v", name, tc.want, got)
				t.Log("失败！")
			}
		})
	}
}
