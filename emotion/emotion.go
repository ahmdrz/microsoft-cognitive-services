package emotion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func New(key string) (*Emotion, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Emotion{
		BingKey: key,
	}, nil
}

func (emotion *Emotion) Recognize(url string) ([]EmotionDetail, error) {
	apiURL := URL + "/recognize"

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", emotion.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := make([]EmotionDetail, 0)
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 429 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(result.Code)
	}

	return nil, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (emotion *Emotion) RecognizeFace(url string, faceArr []Rectangle) ([]EmotionDetail, error) {
	if len(faceArr) < 1 {
		return nil, fmt.Errorf("Incorrect face array")
	}

	faces := ""
	for index, face := range faceArr {
		faces += face.encode()
		if index < len(faceArr)-1 {
			faces += ";"
		}
	}

	apiURL := URL + "/recognize?faceRectangles=" + faces

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", emotion.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := make([]EmotionDetail, 0)
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	if resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 429 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(result.Code)
	}

	return nil, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (emotion *Emotion) GetVideoResult(id string) (VideoResult, error) {
	apiURL := URL + "/operations/" + id

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return VideoResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", emotion.BingKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VideoResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := VideoResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VideoResult{}, err
		}

		return result, nil
	}

	return VideoResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}
