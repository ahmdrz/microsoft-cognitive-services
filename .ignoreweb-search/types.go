package websearch

type WebSearch struct {
	BingKey       string
	LastRequestID string
	Language      string
	SafeSearch    string
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}

const (
	URL                     string = "https://api.cognitive.microsoft.com/bing/v5.0"
	LANG_DEFAULT            string = "en-US"
	LANG_ChineseSimplified  string = "zh-Hans"
	LANG_ChineseTraditional string = "zh-Hant"
	LANG_Czech              string = "cs"
	LANG_Danish             string = "da"
	LANG_Dutch              string = "nl"
	LANG_English            string = "en"
	LANG_Finnish            string = "fi"
	LANG_French             string = "fr"
	LANG_German             string = "de"
	LANG_Greek              string = "el"
	LANG_Hungarian          string = "hu"
	LANG_Italian            string = "it"
	LANG_Japanese           string = "Ja"
	LANG_Korean             string = "ko"
	LANG_Norwegian          string = "nb"
	LANG_Polish             string = "pl"
	LANG_Portuguese         string = "pt"
	LANG_Russian            string = "ru"
	LANG_Spanish            string = "es"
	LANG_Swedish            string = "sv"
	LANG_Turkish            string = "tr"
)

type SearchResult struct {
	Type            string          `json:"_type"`
	Instrumentation Instrumentation `json:"instrumentation"`
	WebPages        WebPages        `json:"webPages"`
	Images          Images          `json:"images"`
	News            News            `json:"news"`
	RelatedSearches RelatedSearches `json:"relatedSearches"`
	Videos          Videos          `json:"videos"`
}
