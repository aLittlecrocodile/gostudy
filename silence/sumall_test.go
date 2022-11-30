package silence

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 9}, []int{1, 1, 1})
	expected := []int{9, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)

	}
}
func TestSumAllTails(t *testing.T) {
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		wanted := []int{6, 9}
		if reflect.DeepEqual(got, wanted) {
			t.Errorf("want %v but got %v", wanted, got)

		}

	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		wanted := []int{0, 9}
		if reflect.DeepEqual(got, wanted) {
			t.Errorf("want %v but got %v", wanted, got)

		}

	})

}
