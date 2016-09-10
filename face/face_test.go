package face

import (
	"fmt"
	"testing"
)

var key string = "-" // this code is sample , forget this :) my token has been changed !

func TestNew(t *testing.T) {
	_,err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
}
