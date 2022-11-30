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
