package rsa

import (
	"testing"
)

var tests = []struct {
	expected int
	arg1     int
	arg2     int
}{
	{2, 2, 6},
	{2, 4, 6},
	{13, 26, 13},
	{13, 13, 26},
}

func Test_greatestCommonDivisor1(t *testing.T) {
	for _, tt := range tests {
		res := greatestCommonDivisor1(tt.arg1, tt.arg2)
		if res != tt.expected {
			t.Errorf("expected:%v but got:%v", tt.expected, res)
		}
	}
}
func Test_greatestCommonDivisor2(t *testing.T) {
	for _, tt := range tests {
		res := greatestCommonDivisor2(tt.arg1, tt.arg2)
		if res != tt.expected {
			t.Errorf("expected:%v but got:%v", tt.expected, res)
		}

	}
}

func Test_greatestCommonDivisor3(t *testing.T) {
	for _, tt := range tests {
		res := greatestCommonDivisor3(tt.arg1, tt.arg2)
		if res != tt.expected {
			t.Errorf("expected:%v but got:%v", tt.expected, res)
		}
	}
}
