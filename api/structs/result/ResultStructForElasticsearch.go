package result

type ResultStructForElasticSearch struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			ID     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
