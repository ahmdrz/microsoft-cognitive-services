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
	result, err := spell.Check("Gates", MODE_SPELL)
	fmt.Println(result)
}
