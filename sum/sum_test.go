package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any numbers", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		
		if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	// t.Run("Collection of 3 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6
		
	// 	if got != want {
	// 	t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	t.Run("add two slice to one slice", func(t *testing.T){
		sumSlice := SumAll([]int{1,2,3}, []int{1,2,3})

		got := sumSlice

		want := []int{6, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}