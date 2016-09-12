package face

import (
	"fmt"
	"reflect"
)

type Face struct {
	BingKey       string
	LastRequestID string
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}

const (
	URL string = "https://api.projectoxford.ai/face/v1.0"
)

type HeadPose struct {
	Roll  float32 `json:"roll"`
	Pitch float32 `json:"pitch"`
	Yaw   float32 `json:"yaw"`
}

type FacialHair struct {
	Mustache  float32 `json:"mustache"`
	Beard     float32 `json:"beard"`
	SideBurns float32 `json:"sideburns"`
}

type FaceAttributes struct {
	Age        float32    `json:"age"`
	Gender     string     `json:"gender"`
	Smile      float32    `json:"smile"`
	FacialHair FacialHair `json:"facialhair"`
	Glasses    string     `json:"glasses"`
	HeadPose   HeadPose   `json:"headpose"`
}

type Rectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FaceLandmarks struct {
	UnderLipBottom      Position `json:"underLipBottom"`
	UnderLipTop         Position `json:"underLipTop"`
	UpperLipBottom      Position `json:"upperLipBottom"`
	UpperLipTop         Position `json:"upperLipTop"`
	NoseRightAlarOutTip Position `json:"noseRightAlarOutTip"`
	NoseLeftAlarOutTip  Position `json:"noseLeftAlarOutTip"`
	NoseRightAlarTop    Position `json:"noseRightAlarTop"`
	NoseLeftAlarTop     Position `json:"noseLeftAlarTop"`
	NoseRootRight       Position `json:"noseRootRight"`
	NoseRootLeft        Position `json:"noseRootLeft"`
	EyeRightOuter       Position `json:"eyeRightOuter"`
	EyeRightBottom      Position `json:"eyeRightBottom"`
	EyeRightTop         Position `json:"eyeRightTop"`
	EyeRightInner       Position `json:"eyeRightInner"`
	EyebrowRightOuter   Position `json:"eyebrowRightOuter"`
	EyebrowRightInner   Position `json:"eyebrowRightInner"`
	EyeLeftInner        Position `json:"eyeLeftInner"`
	EyeLeftBottom       Position `json:"eyeLeftBottom"`
	EyeLeftTop          Position `json:"eyeLeftTop"`
	EyeLeftOuter        Position `json:"eyeLeftOuter"`
	EyebrowLeftInner    Position `json:"eyebrowLeftInner"`
	EyebrowLeftOuter    Position `json:"eyebrowLeftOuter"`
	MouthRight          Position `json:"mouthRight"`
	MouthLeft           Position `json:"mouthLeft"`
	NoseTip             Position `json:"noseTip"`
	PupilRight          Position `json:"pupilRight"`
	PupilLeft           Position `json:"pupilLeft"`
}

type FaceAttributesOrder struct {
	Age        bool `name:"age"`
	Gender     bool `name:"gender"`
	HeadPose   bool `name:"headPose"`
	Smile      bool `name:"smile"`
	FacialHair bool `name:"facialHair"`
	Glasses    bool `name:"glasses"`
}

func (order FaceAttributesOrder) String() (string, error) {
	var (
		v = make([]string, 0)
		s = reflect.ValueOf(order)
		t = reflect.TypeOf(order)
	)

	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface().(bool) == true {
			field := t.Field(i)
			v = append(v, field.Tag.Get("name"))
		}
	}

	if len(v) == 0 {
		return "", fmt.Errorf("Empty v")
	}

	result := ""
	for _, item := range v {
		result += item + ","
	}

	return result[:len(result)-1], nil
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type FaceDetect struct {
	FaceID         string         `json:"faceId"`
	FaceRectangle  Rectangle      `json:"faceRectangle"`
	FaceLandmarks  FaceLandmarks  `json:"faceLandmarks"`
	FaceAttributes FaceAttributes `json:"faceAttributes"`
}

type DetectOrder struct {
	FaceID         bool
	FaceLandmarks  bool
	FaceAttributes FaceAttributesOrder
}

func (order DetectOrder) String() string {
	boolToString := func(i bool) string {
		if i {
			return "true"
		}
		return "false"
	}
	str, err := order.FaceAttributes.String()
	if err != nil {
		return "returnFaceId=" + boolToString(order.FaceID) + "&returnFaceLandmarks=" + boolToString(order.FaceLandmarks)
	}
	return "returnFaceId=" + boolToString(order.FaceID) + "&returnFaceLandmarks=" + boolToString(order.FaceLandmarks) + "&returnFaceAttributes=" + str
}
