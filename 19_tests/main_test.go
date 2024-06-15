package main

import (
	"testing"
)

var tests = []struct {
	name    string
	a       float32
	b       float32
	want    float32
	isError bool
}{
	{"valide", 10.0, 2.0, 5.0, false},
	{"invalid", 10.0, 0.0, 0.0, true},
	{"valide", -10.0, 2.0, -5.0, false},
	{"valide", -10.0, -2.0, 5.0, false},
	{"valide", 10.0, -2.0, -5.0, false},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.a, tt.b)

			if (err != nil) != tt.isError {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.isError)
				return
			}

			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
