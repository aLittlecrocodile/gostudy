package count

import (
	"fmt"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	expect := 2
	got := strings.Count("apple", "p")
	if expect != got {
		t.Errorf("expected %d but got $d", expect, got)
	}

}
func ExampleCount() {
	expected := strings.Count("aaa", "a")
	fmt.Println(expected)
	// Output: 3
}
