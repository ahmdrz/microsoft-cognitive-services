// vision project vision.go
// license MIT powered by github.com/ahmdrz
// full api reference on https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/
package vision

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func toStringBool(a bool) string {
	if a {
		return "true"
	}
	return "false"
}

func toString(a int) string {
	return strconv.Itoa(a)
}

func New(key string) (*Vision, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Vision{
		BingKey: key,
	}, nil
}

func (vision *Vision) Analyze(url string, order VisualFeatures) (VisionResult, error) {
	query, err := order.String()
	if err != nil {
		panic(err)
	}
	apiURL := URL + "/analyze?visualFeatures=" + query

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := VisionResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		vision.LastRequestID = result.RequestID
		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		return VisionResult{}, fmt.Errorf(result.Code)
	}

	return VisionResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (vision *Vision) Tag(url string) (VisionResult, error) {
	apiURL := URL + "/tag"

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := VisionResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		vision.LastRequestID = result.RequestID
		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		return VisionResult{}, fmt.Errorf(result.Code)
	}

	return VisionResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (vision *Vision) OCR(url string, order OCROption) (VisionOCRResult, error) {
	query := order.String()
	apiURL := URL + "/ocr?" + query

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return VisionOCRResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionOCRResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := VisionOCRResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionOCRResult{}, err
		}

		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionOCRResult{}, err
		}

		return VisionOCRResult{}, fmt.Errorf(result.Code)
	}

	return VisionOCRResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (vision *Vision) Describe(url string, max int) (VisionResult, error) {
	apiURL := URL + "/describe?maxCandidates=" + toString(max)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := VisionResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		vision.LastRequestID = result.RequestID
		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		return VisionResult{}, fmt.Errorf(result.Code)
	}

	return VisionResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}

func (vision *Vision) GetModels() ([]Model, error) {
	apiURL := URL + "/models"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		var result struct {
			Models []Model `json:"models"`
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return result.Models, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
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

func (vision *Vision) Thumbnail(url string, order ThumbnailOrder) (Image, error) {
	if order.Width < 1 || order.Width > 1024 {
		return Image{}, fmt.Errorf("Incorrect width")
	}
	if order.Height < 1 || order.Height > 1024 {
		return Image{}, fmt.Errorf("Incorrect height")
	}

	apiURL := URL + "/generateThumbnail?width=" + toString(order.Width) + "&height=" + toString(order.Height) + "&smartCropping=" + toStringBool(order.SmartCropping)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return Image{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Image{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		return Image{
			Image: body,
		}, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return Image{}, err
		}

		return Image{}, fmt.Errorf(result.Code)
	}

	return Image{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}
