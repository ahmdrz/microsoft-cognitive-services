package spell

type Spell struct {
	BingKey       string
	LastRequestID string
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}

const (
	URL        string = "https://api.cognitive.microsoft.com/bing/v5.0"
	MODE_SPELL string = "spell"
	MODE_PROOF string = "proof"
)

type Result struct {
	Type          string         `json:"_type"`
	FlaggedTokens []FlaggedToken `json:"flaggedTokens"`
}

type FlaggedToken struct {
	Offset      int       `json:"offset"`
	Token       string    `json:"token"`
	Type        string    `json:"type"`
	Suggestions []Suggest `json:"suggestions"`
}

type Suggest struct {
	Suggestion string `json:"suggestion"`
	Score      int    `json:"score"`
}
