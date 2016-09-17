package spell

import (
	"fmt"
	"testing"
)

var key string = "-"

func TestCheck(t *testing.T) {
	spell, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := spell.Check("Richard Stalman", MODE_SPELL)
	fmt.Println(result, err)
}
