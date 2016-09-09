package vision

import (
	"fmt"
	"testing"
)

var key string = "4f558a77503b46549ac0d784c1651d9e" // this code is sample , forget this :) my token has been changed !

var testFile string = "test.jpg"

func TestThumbnail(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Thumbnail("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", ThumbnailOrder{
		Width:  10,
		Height: 10,
	})
	result.Save("temp.jpg")
	if err != nil {
		t.Log(err)
		return
	}
}

func TestGetModels(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.GetModels()
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestFileOCR(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.OCRFile(testFile, OCROption{
		Language: LANG_English,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result.String())
}

func TestFileDescribe(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.DescribeFile(testFile, 1)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestFileTag(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.TagFile(testFile)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestDescribe(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Describe("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", 1)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestFileAnalyze(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.AnalyzeFile(testFile, VisualFeatures{
		Tags: true,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestOCR(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.OCR("https://portalstoragewuprod2.azureedge.net/vision/OpticalCharacterRecognition/1.jpg", OCROption{
		Language: LANG_English,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result.String())
}

func TestAnalyze(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", VisualFeatures{
		Tags: true,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestTag(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Tag("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestMakeString(t *testing.T) {
	vis := VisualFeatures{}
	vis.Adult = true
	vis.Categories = true

	fmt.Println(vis.String())
}
