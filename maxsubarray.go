package algorithm

func FindMaximumSubarrayLinear(in []int, low, high int) (left, right, sum int) {
	left = 0
	right = 0
	sum = in[low]
	temp_sum := 0
	temp_left := 0
	for i := low; i <= high; i++ {
		if in[i] > (temp_sum + in[i]) {
			temp_sum = in[i]
		} else {
			temp_sum = temp_sum + in[i]
		}
		if temp_sum == in[i] {
			temp_left = i
		}
		if temp_sum > sum {
			sum = temp_sum
			right = i
			left = temp_left
		}
	}
	return left, right, sum
}

func findMaxCrossingSubarray(in []int, low, mid, high int) (left, right, maxSum int) {
	leftSum := MinInt
	sum := 0
	for i := mid; i >= low; i-- {
		sum = sum + in[i]
		if sum > leftSum {
			leftSum = sum
			left = i
		}
	}
	rightSum := MinInt
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum = sum + in[j]
		if sum > rightSum {
			rightSum = sum
			right = j
		}
	}
	return left, right, leftSum + rightSum
}

func FindMaximumSubarray(in []int, low, high int) (left, right, sum int) {
	if high == low {
		return low, high, in[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaximumSubarray(in, low, mid)
		rightLow, rightHigh, rightSum := FindMaximumSubarray(in, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubarray(in, low, mid, high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}
