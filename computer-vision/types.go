package vision

import (
	"fmt"
	"os"
	"reflect"
)

const (
	URL                     string = "https://api.projectoxford.ai/vision/v1.0"
	LANG_AutoDetect         string = "unk"
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

// Main struct of library , contains BingKey for saving token and LastRequestID for determine the request id of last method.
type Vision struct {
	BingKey       string
	LastRequestID string
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

/*
Categories - categorizes image content according to a taxonomy defined in documentation.

Tags - tags the image with a detailed list of words related to the image content.

Description - describes the image content with a complete English sentence.

Faces - detects if faces are present. If present, generate coordinates, gender and age.

ImageType - detects if image is clipart or a line drawing.

Color - determines the accent color, dominant color, and whether an image is black&white.

Adult - detects if the image is pornographic in nature (depicts nudity or a sex act). Sexually suggestive content is also detected.
*/

type VisualFeatures struct {
	Categories  bool
	Tags        bool
	Description bool
	Faces       bool
	ImageType   bool
	Color       bool
	Adult       bool
}

func (order VisualFeatures) String() (string, error) {
	var (
		v = make([]string, 0)
		s = reflect.ValueOf(order)
		t = reflect.TypeOf(order)
	)

	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface().(bool) == true {
			v = append(v, t.Field(i).Name)
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

type OCROption struct {
	Language          string
	DetectOrientation bool
}

func (order OCROption) String() string {
	if order.Language == "" {
		order.Language = LANG_AutoDetect
	}
	if order.DetectOrientation {
		return "language=" + order.Language + "&detectOrientation=" + "true"
	}
	return "language=" + order.Language + "&detectOrientation=" + "false"
}

type VisionOCRResult struct {
	Language    string   `json:"language"`
	TextAngle   float64  `json:"textAngle"`
	Orientation string   `json:"orientation"`
	Regions     []Region `json:"regions"`
}

// this method can make a full sentence from OCR result
func (order VisionOCRResult) String() (result string) {
	result = ""
	for _, region := range order.Regions {
		for _, line := range region.Lines {
			for index, word := range line.Words {
				result += word.Text
				if index < len(line.Words)-1 {
					result += " "
				}
			}
			result += "\n"
		}
	}
	return
}

type Region struct {
	BoundingBox string `json:"boundingBox"`
	Lines       []Line `json:"lines"`
}

type Line struct {
	BoundingBox string `json:"boundingBox"`
	Words       []Word `json:"words"`
}

type Word struct {
	BoundingBox string `json:"boundingBox"`
	Text        string `json:"text"`
}

type VisionResult struct {
	RequestID   string      `json:"requestId"`
	Categories  []Category  `json:"categories"`
	Adult       Adult       `json:"adult"`
	Faces       []Face      `json:"faces"`
	ImageType   ImageType   `json:"imageType"`
	Color       Color       `json:"color"`
	MetaData    MetaData    `json:"metadata"`
	Description Description `json:"description"`
	Tags        []Tag       `json:"tags"`
}

type Tag struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

type Description struct {
	Tags     []string  `json:"tags"`
	Captions []Caption `json:"captions"`
}

type Caption struct {
	Text       string  `json:"text"`
	Confidence float64 `json:"confidence"`
}

type MetaData struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}

type Color struct {
	DominantColorForeground string   `json:"dominantColorForeground"`
	DominantColorBackground string   `json:"dominantColorBackground"`
	AccentColor             string   `json:"accentColor"`
	IsBWImg                 bool     `json:"isBWImg"`
	DominantColors          []string `json:"dominantColors"`
}

type ImageType struct {
	ClipArtType     int `json:"clipArtType"`
	LineDrawingType int `json:"lineDrawingType"`
}

type Face struct {
	Age           int       `json:"age"`
	Gender        string    `json:"gender"`
	FaceRectangle Rectangle `json:"faceRectangle"`
}

type Rectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Adult struct {
	IsAdultContent bool    `json:"isAdultContent"`
	IsRacyContent  bool    `json:"isRacyContent"`
	AdultScore     float64 `json:"adultScore"`
	RacyScore      float64 `json:"racyScore"`
}

type Category struct {
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}

type Model struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type Image struct {
	Image []byte
}

func (image Image) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(image.Image)
	if err != nil {
		return err
	}
	return f.Sync()
}

type ThumbnailOrder struct {
	Width         int
	Height        int
	SmartCropping bool
}
