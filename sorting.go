package algorithm

func InsertionSort(in []int) {
	for j := 1; j < len(in); j++ {
		key := in[j]
		i := j - 1
		for i >= 0 && in[i] > key {
			in[i+1] = in[i]
			i--
		}
		in[i+1] = key
	}
}

func merge(in []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1)
	right := make([]int, n2)
	for i := 0; i < n1; i++ {
		left[i] = in[p+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = in[q+1+j]
	}
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if j >= n2 {
			in[k] = left[i]
			i++
		} else if i >= n1 {
			in[k] = right[j]
			j++
		} else {
			if left[i] <= right[j] {
				in[k] = left[i]
				i++
			} else {
				in[k] = right[j]
				j++
			}
		}
	}
}

func MergeSort(in []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(in, p, q)
		MergeSort(in, q+1, r)
		merge(in, p, q, r)
	}
}
