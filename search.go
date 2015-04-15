package elasticsearchgo

import (
	"encoding/json"
	"github.com/magicshui/goutils/requests"
)

func (this *EsClient) Search() (Response, error) {

	var response Response
	url := this.Domain + "/" + this.Index + "/" + this.Type + "/_search"
	query := this.Query

	data, err := requests.PostRequest(url, query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	_Shards  struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int           `json:"total"`
		MaxScore float32       `json:"max_score"`
		Hits     []interface{} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		Groupbyevent struct {
			DocCountErrorUpperBound int           `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int           `json:"sum_other_doc_count"`
			Buckets                 []interface{} `json:"buckets"`
		} `json:"groupbyevent"`
	} `json:"aggregations"`
}
