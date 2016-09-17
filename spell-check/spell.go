package spell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func New(key string) (*Spell, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Spell{
		BingKey: key,
	}, nil
}

func (spell *Spell) Check(word string, mode string) (Result, error) {
	apiURL := URL + "/spellcheck?mode=" + mode
	data := url.Values{}
	data.Set("Text", word)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return Result{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", spell.BingKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := Result{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return Result{}, err
		}

		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 429 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return Result{}, err
		}

		return Result{}, fmt.Errorf(result.Code)
	}

	return Result{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}
