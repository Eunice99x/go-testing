package main

import "fmt"



func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) != 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
			fmt.Println("############## IF ############")
			fmt.Println("=========> len(sums)", len(sums))
			fmt.Println("=========> sums", sums)
			fmt.Println("=========>  len(numbers)", len(numbers))
			fmt.Println("=========>hhhhhhhhhhhhhhhh  len(numbers)", numbers)
		} else {
			fmt.Println("############## ELSE ############")
			fmt.Println("=========> len(sums)", len(sums))
			fmt.Println("=========> sums", sums)
			sums = append(sums, 0)
		}
		fmt.Println("=========> numbersToSum", numbersToSum)
		fmt.Println("=========> len(numbersToSum)", len(numbersToSum))
	}
	
	fmt.Println("############## RES ############")
	fmt.Println("=========> len(sums)", len(sums))
	fmt.Println("=========> sums", sums)
	return sums
}
func main() {
	SumAllTails([]int{1, 2, 3}, []int{})
}