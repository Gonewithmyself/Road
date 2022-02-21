package main

import "testing"

func Test_testarray(t *testing.T) {
	type args struct {
		arr []interface{}
		a   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{arr: []interface{}{1, 2, myint(3)}, a: 3}, true},
		{"", args{arr: []interface{}{1, 2, uint16(3)}, a: 3}, true},
		{"", args{arr: []interface{}{1, 2, 3}, a: 3}, true},
		{"", args{arr: []interface{}{1, 2, 3}, a: 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testarray(tt.args.arr, tt.args.a); got != tt.want {
				t.Errorf("testarray() = %v, want %v", got, tt.want)
			}
		})
	}
	// 	t.Error()
}
