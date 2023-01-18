package main

import "testing"

//// TestAdd with testify
//func TestAdd(t *testing.T) {
//	a := assert.New(t)
//	a.Equal(Add(1,2), 3)
//	a.Equal(Add(10,20), 99)	// error
//}


// TestAdd auto generate
func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", struct {
			a int
			b int
		}{a: 1, b: 2},
			3,
		},
		{
			"t2", struct {
			a int
			b int
		}{a: 10, b: 20},
			30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
