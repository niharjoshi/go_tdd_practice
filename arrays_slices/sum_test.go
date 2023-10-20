package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Testing sum of multiple collections", func(t *testing.T) {
		numbers_1 := []int{1, 2, 3, 4}
		numbers_2 := []int{5, 6, 7, 8, 9, 10}
		got := SumAll(numbers_1, numbers_2)
		want := []int{10, 45}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("Testing sum of tails of multiple collections", func(t *testing.T) {
		numbers_1 := []int{1, 2, 3, 4}
		numbers_2 := []int{5, 6, 7, 8, 9, 10}
		got := SumAllTails(numbers_1, numbers_2)
		want := []int{9, 40}
		checkSums(t, got, want)
	})
	t.Run("Testing sum of empty collections", func(t *testing.T) {
		numbers_1 := []int{}
		numbers_2 := []int{5, 6, 7, 8, 9, 10}
		got := SumAllTails(numbers_1, numbers_2)
		want := []int{0, 40}
		checkSums(t, got, want)
	})
}
