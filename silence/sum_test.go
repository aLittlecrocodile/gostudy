package silence

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		wanted := 15
		if got != wanted {
			t.Errorf("wanted %d but got %d,%v", wanted, got, numbers)
		}

	})
	t.Run("collection of any numbers ", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		wanted := 6
		if got != wanted {
			t.Errorf("wanted %d but got %d,%v", wanted, got, numbers)
		}

	})

}
