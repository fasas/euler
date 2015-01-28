// Copyright 2014 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import "testing"

var isPrimeTests = []struct {
	in  int
	out bool
}{
	{-4, false},
	{-3, false},
	{-2, false},
	{-1, false},
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{6, false},
	{7, true},
	{11, true},
	{12, false},
	{15, false},
	{17, true},
	{23, true},
	{24, false},
	{25, false},
	{29, true},
	{31, true},
	{100, false},
	{105, false},
}

func TestIsPrime(t *testing.T) {
	for _, tt := range isPrimeTests {
		if IsPrime(tt.in) != tt.out {
			t.Errorf("IsPrime(%d) = %t; want %t", tt.in, !tt.out, tt.out)
		}
	}
}

var mulTests = []struct {
	a   int
	b   int
	out int
	ok  bool
}{
	{MinInt, 0, 0, true},
	{0, MinInt, 0, true},
	{MinInt, 1, MinInt, true},
	{1, MinInt, MinInt, true},
	{MaxInt, -1, -MaxInt, true},
	{-1, MaxInt, -MaxInt, true},
	{-MaxInt, -1, MaxInt, true},
	{-1, -MaxInt, MaxInt, true},
	{MinInt, 2, 0, false},
	{MinInt, -2, 0, false},
	{MinInt, -1, 0, false},
	{2, MinInt, 0, false},
	{-2, MinInt, 0, false},
	{-1, MinInt, 0, false},
}

func TestMul(t *testing.T) {
	for _, tt := range mulTests {
		got, ok := Mul(tt.a, tt.b)
		if got != tt.out || ok != tt.ok {
			t.Errorf("Mul(%d, %d) = %d, %t; want %d, %t", tt.a, tt.b, got, ok, tt.out, tt.ok)
		}
	}
}

var mulIntsTests = []struct {
	in  []int
	out int
	ok  bool
}{
	{[]int{-2, 3, 4}, -24, true},
	{[]int{-3, -4, 5}, 60, true},
	{[]int{-4, -5, -6}, -120, true},
	{[]int{0, 0, 0}, 0, true},
	{[]int{2, 3, 4}, 24, true},
	{[]int{3, 4, 5}, 60, true},
	{[]int{4, 5, 6}, 120, true},
	{[]int{MinInt, 2}, 0, false},
	{[]int{MinInt, -2}, 0, false},
	{[]int{MinInt, -1}, 0, false},
	{[]int{2, MinInt}, 0, false},
	{[]int{-2, MinInt}, 0, false},
	{[]int{-1, MinInt}, 0, false},
}

func TestMulInts(t *testing.T) {
	for _, tt := range mulIntsTests {
		got, ok := MulInts(tt.in)
		if got != tt.out || ok != tt.ok {
			t.Errorf("MulInts(%d) = %d, %t; want %d, %t", tt.in, got, ok, tt.out, tt.ok)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(i)
	}
}

func BenchmarkMulInts(b *testing.B) {
	ints := make([]int, 1e3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MulInts(ints)
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(i, i)
	}
}