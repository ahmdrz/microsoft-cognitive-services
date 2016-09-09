package emotion

import (
	_ "encoding/json"
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

func (emotion *Emotion) Recognize(url string) {
	apiURL := URL + "/recognize"

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", emotion.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
