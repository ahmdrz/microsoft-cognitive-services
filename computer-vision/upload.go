package vision

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func (vision *Vision) AnalyzeFile(path string, order VisualFeatures) (VisionResult, error) {
	query, err := order.String()
	if err != nil {
		return VisionResult{}, err
	}
	apiURL := URL + "/analyze?visualFeatures=" + query
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(path)
	if err != nil {
		return VisionResult{}, err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image", path)
	if err != nil {
		return VisionResult{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return VisionResult{}, err
	}
	w.Close()
	req, err := http.NewRequest("POST", apiURL, &b)
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", w.FormDataContentType())
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

func (vision *Vision) TagFile(path string) (VisionResult, error) {
	apiURL := URL + "/tag"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(path)
	if err != nil {
		return VisionResult{}, err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image", path)
	if err != nil {
		return VisionResult{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return VisionResult{}, err
	}
	w.Close()
	req, err := http.NewRequest("POST", apiURL, &b)
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

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

func (vision *Vision) DescribeFile(path string, max int) (VisionResult, error) {
	apiURL := URL + "/describe?maxCandidates=" + toString(max)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(path)
	if err != nil {
		return VisionResult{}, err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image", path)
	if err != nil {
		return VisionResult{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return VisionResult{}, err
	}
	w.Close()
	req, err := http.NewRequest("POST", apiURL, &b)
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

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

func (vision *Vision) OCRFile(path string, order OCROption) (VisionOCRResult, error) {
	query := order.String()
	apiURL := URL + "/ocr?" + query

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(path)
	if err != nil {
		return VisionOCRResult{}, err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image", path)
	if err != nil {
		return VisionOCRResult{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return VisionOCRResult{}, err
	}
	w.Close()
	req, err := http.NewRequest("POST", apiURL, &b)
	if err != nil {
		return VisionOCRResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", w.FormDataContentType())

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
