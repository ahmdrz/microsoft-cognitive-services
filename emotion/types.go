package emotion

type Emotion struct {
	BingKey       string
	LastRequestID string
}

const (
	URL string = "https://api.projectoxford.ai/emotion/v1.0"
)
