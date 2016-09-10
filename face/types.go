package face

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
