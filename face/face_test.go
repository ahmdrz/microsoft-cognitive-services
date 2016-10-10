package face

import (
	"fmt"
	"testing"
)

var key string = "-" // this code is sample , forget this :) my token has been changed !

func TestOrder(t *testing.T) {
	test := FaceAttributesOrder{
		Age:    true,
		Gender: true,
	}
	fmt.Println(test.String())
}

func TestNew(t *testing.T) {
	_, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
}

func TestDetect(t *testing.T) {
	face, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := face.Detect("https://portalstoragewuprod.azureedge.net/media/Default/Documentation/Face/Images/FaceFindSimilar.QueryFace.jpg",
		DetectOrder{
			FaceLandmarks: true,
			FaceAttributes: FaceAttributesOrder{
				Age: true,
			},
		},
	)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}
