package emotion

import (
	"strconv"
)

type Emotion struct {
	BingKey       string
	LastRequestID string
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}

const (
	URL string = "https://api.projectoxford.ai/emotion/v1.0"
)

type EmotionDetail struct {
	FaceRectangle Rectangle `json:"faceRectangle"`
	Scores        Score     `json:"scores"`
}

type Rectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (rec *Rectangle) encode() string {
	toString := func(a int) string {
		return strconv.Itoa(a)
	}
	return toString(rec.Left) + "," + toString(rec.Top) + "," + toString(rec.Width) + "," + toString(rec.Height)
}

type Score struct {
	Anger     float64 `json:"anger"`
	Contempt  float64 `json:"contempt"`
	Disgust   float64 `json:"disgust"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
}

type VideoResult struct {
	Status             string `json:"status"`
	CreatedDateTime    string `json:"createdDateTime"`
	LastActionDateTime string `json:"lastActionDateTime"`
}
