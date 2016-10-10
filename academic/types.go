package academic

import (
	"strconv"
)

type Academic struct {
	BingKey       string
	LastRequestID string
}

type Error struct {
	Code    string `json:"statusCode"`
	Message string `json:"message"`
}

const (
	URL string = "https://api.projectoxford.ai/academic/v1.0"
)

type CalcHistogramRequest struct {
	Model      string
	Attributes string
	Count      int
	Offset     int
}

func (c CalcHistogramRequest) String() string {
	result := ""
	if len(c.Model) == 0 {
		c.Model = "latest"
	}

	result += "&model=" + c.Model

	if len(c.Attributes) != 0 {
		result += "&model=" + c.Attributes
	}

	result += "&count=" + strconv.Itoa(c.Count)
	result += "&offset=" + strconv.Itoa(c.Offset)

	return result
}

type CalcHistogramResult struct {
	Expr        string          `json:"expr"`
	NumEntities int             `json:"num_entities"`
	Histograms  []HistogramItem `json:"histograms"`
}

type HistogramItem struct {
	Attribute      string      `json:"attribute"`
	DistinctValues int         `json:"distinct_values"`
	TotalCount     int         `json:"total_count"`
	Histogram      []Histogram `json:"histogram"`
}

type Histogram struct {
	Value int     `json:"value"`
	Prob  float64 `json:"prob"`
	Count int     `json:"count"`
}
