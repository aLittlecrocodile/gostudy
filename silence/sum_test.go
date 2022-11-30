package silence

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	wanted := 15
	if got != wanted {
		t.Errorf("wanted %d but got %d,%v", wanted, got, numbers)
	}
}
