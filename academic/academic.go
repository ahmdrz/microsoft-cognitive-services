package academic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func New(key string) (*Academic, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Academic{
		BingKey: key,
	}, nil
}

func (academic *Academic) CalcHistogram(expr string, order CalcHistogramRequest) (CalcHistogramResult, error) {
	values := url.Values{}
	values.Set("expr", expr)
	if len(order.Model) == 0 {
		order.Model = "latest"
	}
	values.Set("model", order.Model)
	if len(order.Attributes) != 0 {
		values.Set("attributes", order.Attributes)
	}
	values.Set("count", strconv.Itoa(order.Count))
	values.Set("offset", strconv.Itoa(order.Offset))
	apiURL := URL + "/calchistogram?" + values.Encode()
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return CalcHistogramResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", academic.BingKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CalcHistogramResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result := CalcHistogramResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return CalcHistogramResult{}, err
		}

		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 429 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return CalcHistogramResult{}, err
		}

		return CalcHistogramResult{}, fmt.Errorf(result.Code)
	}

	return CalcHistogramResult{}, fmt.Errorf("Unknown Error Occurred , Check the key , Status : " + resp.Status)
}
