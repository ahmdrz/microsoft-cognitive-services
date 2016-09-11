package face

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func New(key string) (*Face, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Face{
		BingKey: key,
	}, nil
}

func (face *Face) Detect(url string, order DetectOrder) ([]FaceDetect, error) {
	apiURL := URL + "/detect?" + order.String()
	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", face.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := make([]FaceDetect, 0)
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
