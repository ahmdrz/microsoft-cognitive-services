package emotion

import (
	_ "fmt"
	"testing"
)

var key string = "-" // this code is sample , forget this :) my token has been changed !

func TestRecognize(t *testing.T) {
	emotion, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	emotion.Recognize("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
}
