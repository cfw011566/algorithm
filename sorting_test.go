package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleInsertionSort() {
	in := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(in)
	fmt.Println(in)
	//!-main

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

func benchmarkInsertionSort(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		in := make([]int, size)
		for j := 0; j < len(in); j++ {
			in[j] = rand.Int()
		}
		InsertionSort(in)
	}
}

func BenchmarkInsertionSort100(b *testing.B) {
	benchmarkInsertionSort(b, 100)
}

func BenchmarkInsertionSort1000(b *testing.B) {
	benchmarkInsertionSort(b, 1000)
}

func BenchmarkInsertionSort10000(b *testing.B) {
	benchmarkInsertionSort(b, 10000)
}

/*
func BenchmarkInsertionSortWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := make([]int, 10000)
		for j := 0; j < len(in); j++ {
			in[j] = len(in) - j
		}
		InsertionSort(in)
	}
}

func BenchmarkInsertionSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := make([]int, 10000)
		for j := 0; j < len(in); j++ {
			in[j] = j
		}
		InsertionSort(in)
	}
}
*/

func ExampleMergeSort() {
	in := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(in, 0, len(in)-1)
	fmt.Println(in)
	//!-main

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

func benchmarkMergeSort(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		in := make([]int, size)
		for j := 0; j < len(in); j++ {
			in[j] = rand.Int()
		}
		MergeSort(in, 0, len(in)-1)
	}
}

func BenchmarkMergeSort100(b *testing.B) {
	benchmarkMergeSort(b, 100)
}

func BenchmarkMergeSort1000(b *testing.B) {
	benchmarkMergeSort(b, 1000)
}

func BenchmarkMergeSort10000(b *testing.B) {
	benchmarkMergeSort(b, 10000)
}

/*
func BenchmarkMergeSortWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := make([]int, 10000)
		for j := 0; j < len(in); j++ {
			in[j] = len(in) - j
		}
		MergeSort(in, 0, len(in)-1)
	}
}

func BenchmarkMergeSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := make([]int, 10000)
		for j := 0; j < len(in); j++ {
			in[j] = j
		}
		MergeSort(in, 0, len(in)-1)
	}
}
*/
