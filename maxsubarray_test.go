package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleFindMaximumSubarray() {
	in := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(FindMaximumSubarray(in, 0, len(in)-1))
	in = []int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 20, -215}
	fmt.Println(FindMaximumSubarrayLinear(in, 0, len(in)-1))
	in = []int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 20, -2, 30, -1}
	fmt.Println(FindMaximumSubarrayLinear(in, 0, len(in)-1))
	in = []int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 10, 10, -2, 15, 15, -215}
	fmt.Println(FindMaximumSubarray(in, 0, len(in)-1))
	in = []int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 200, -215}
	fmt.Println(FindMaximumSubarray(in, 0, len(in)-1))
	//!-main

	// Output:
	// 7 10 43
	// 2 9 26
	// 11 13 48
	// 11 15 48
	// 11 11 200
}

func TestFindMaximumSubarray(t *testing.T) {
	var tests = []struct {
		input            []int
		left, right, sum int
	}{
		{[]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}, 7, 10, 43},
		{[]int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 20, -215}, 2, 9, 26},
		{[]int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 20, -2, 30, -1}, 11, 13, 48},
		{[]int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 10, 10, -2, 15, 15, -215}, 11, 15, 48},
		{[]int{0, -1, 7, 6, -5, 12, -3, -2, 10, 1, -50, 200, -215}, 11, 11, 200},
	}
	for _, test := range tests {
		left, right, sum := FindMaximumSubarray(test.input, 0, len(test.input)-1)
		if left != test.left || right != test.right || sum != test.sum {
			t.Errorf("FindMaximumSubarray(%v) Failed", test.input)
		}
		left, right, sum = FindMaximumSubarrayLinear(test.input, 0, len(test.input)-1)
		if left != test.left || right != test.right || sum != test.sum {
			t.Errorf("FindMaximumSubarrayLinear(%v) Failed", test.input)
		}
	}
}

func benchmarkMaximumSubarrayLinear(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		in := make([]int, size)
		for j := 0; j < len(in); j++ {
			in[j] = rand.Int()%200 - 100
		}
		FindMaximumSubarrayLinear(in, 0, len(in)-1)
	}
}

func BenchmarkMaximumSubarrayLinear100(b *testing.B) {
	benchmarkMaximumSubarrayLinear(b, 100)
}

func BenchmarkMaximumSubarrayLinear1000(b *testing.B) {
	benchmarkMaximumSubarrayLinear(b, 1000)
}

func BenchmarkMaximumSubarrayLinear10000(b *testing.B) {
	benchmarkMaximumSubarrayLinear(b, 10000)
}

func benchmarkMaximumSubarray(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		in := make([]int, size)
		for j := 0; j < len(in); j++ {
			in[j] = rand.Int()%200 - 100
		}
		FindMaximumSubarray(in, 0, len(in)-1)
	}
}

func BenchmarkMaximumSubarray100(b *testing.B) {
	benchmarkMaximumSubarray(b, 100)
}

func BenchmarkMaximumSubarray1000(b *testing.B) {
	benchmarkMaximumSubarray(b, 1000)
}

func BenchmarkMaximumSubarray10000(b *testing.B) {
	benchmarkMaximumSubarray(b, 10000)
}
