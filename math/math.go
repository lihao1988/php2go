package math

import (
	"math"
	"strconv"
)

// Abs absolute value
// php abs
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Round number round
// php round
func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

// Floor returns the greatest integer value less than or equal to x
// php floor
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Ceil returns the least integer value greater than or equal to x
// php ceil
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Max returns the max num in nums
// php max
func Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}

	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		maxNum = math.Max(maxNum, nums[i])
	}

	return maxNum
}

// Min returns the min num in nums
// php min
func Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}

	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		minNum = math.Min(minNum, nums[i])
	}

	return minNum
}

// DecBin Decimal to binary
// php decbin
func DecBin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// DecHex Decimal to hexadecimal
// php dechex
func DecHex(number int64) string {
	return strconv.FormatInt(number, 16)
}
