package sum

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(arr ...[]int) []int {
	lenArr := len(arr)

	res := make([]int, lenArr)

	for i, numbers := range arr {
		res[i] += Sum(numbers)
	}

	return res
}

// func SumAllTails(arr1, arr2 []int) []int {
// 	var (
// 		res  []int
// 		sum1 []int
// 		sum2 []int
// 	)
// 	for i := range arr1 {
// 		if arr1[0] != arr1[i] {
// 			sum1 = append(sum1, arr1[i])
// 		}
// 	}
// 	res = append(res, sum1...)
// 	for i := range arr2 {
// 		if arr2[i] != arr2[0] {
// 			sum2 = append(sum2, arr2[i])
// 		}
// 	}
// 	res = append(res, sum2...)
// 	return res
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) != 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}