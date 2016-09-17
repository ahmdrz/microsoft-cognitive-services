package websearch

func New(key string) (*WebSearch, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &WebSearch{
		BingKey: key,
	}, nil
}

func (websearch *WebSearch) Setting() {

}

func toString(a int) string {
	return strconv.Itoa(a)
}

func (websearch *WebSearch) Search(string query, count, offset int, language, safeSearch string) (SearchResult, error) {
	apiURL := URL + "/search?q=" + query + "&count=" + toString(count) + "&offset=" + toString(offset) + "&mkt=" + language + "&safesearch=" + safeSearch
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", websearch.BingKey)

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
