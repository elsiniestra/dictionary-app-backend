package customtypes

type SearchAPIResponse struct {
	Metadata struct {
		Limit          string `json:"limit"`
		Offset         string `json:"offset"`
		Operation      string `json:"operation"`
		Provider       string `json:"provider"`
		Schema         string `json:"schema"`
		SourceLanguage string `json:"sourceLanguage"`
		Total          string `json:"total"`
	} `json:"metadata"`
	Results []struct {
		ID          string `json:"id"`
		Label       string `json:"label"`
		MatchString string `json:"matchString"`
		MatchType   string `json:"matchType"`
		Region      string `json:"region"`
		Word        string `json:"word"`
	} `json:"results"`
}

func (s SearchAPIResponse) GetWordsList() []string {
	result := make([]string, len(s.Results))

	for i, item := range s.Results {
		result[i] = item.ID
	}

	return result
}
