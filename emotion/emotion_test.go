package emotion

import (
	"fmt"
	"testing"
)

var key string = "-" // this code is sample , forget this :) my token has been changed !

func TestRecognizeFace(t *testing.T) {
	emotion, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := emotion.RecognizeFace("https://portalstoragewuprod.azureedge.net/emotion/recognition1.jpg", []Rectangle{Rectangle{263, 488, 147, 147}})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestRecognize(t *testing.T) {
	emotion, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := emotion.Recognize("https://portalstoragewuprod.azureedge.net/emotion/recognition1.jpg")
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}
