package contentinsights

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	types "github.com/Leapforce-nl/go_types"
)

const (
	apiName string = "ContentInsights"
	apiURL  string = "https://api.contentinsights.com/v2"
)

// ContentInsights stores ContentInsights configuration
//
type ContentInsights struct {
	_domainID int
	_apiKey   string
	_isLive   bool
}

// Response represents highest level of ContentInsights api response
//
type Response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Message *string         `json:"message"`
	Paging  Paging          `json:"paging"`
}

// Paging returns paging info
//
type Paging struct {
	Page  int     `json:"page"`
	Pages int     `json:"pages"`
	Total int     `json:"total"`
	First string  `json:"first"`
	Last  string  `json:"last"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
}

// NewContentInsights initializes ContentInsights client
//
func NewContentInsights(domainID int, apiKey string, isLive bool) (*ContentInsights, error) {
	if domainID == 0 {
		return nil, &types.ErrorString{"Informer ApiUrl not provided"}
	}
	if apiKey == "" {
		return nil, &types.ErrorString{"ApiKey not provided"}
	}

	ci := ContentInsights{}
	ci._domainID = domainID
	ci._apiKey = apiKey
	ci._isLive = isLive

	return &ci, nil
}

// Get is a generic Get method
//
func (ci *ContentInsights) Get(path string, queryParams *url.Values, model interface{}) (*Paging, error) {

	queryParams.Set("domain_id", strconv.Itoa(ci._domainID))
	queryParams.Set("api_key", ci._apiKey)

	url := fmt.Sprintf("%s/%s?%s", apiURL, path, queryParams.Encode())

	return ci.GetURL(url, model)
}

// GetURL is a generic Get method
//
func (ci *ContentInsights) GetURL(url string, model interface{}) (*Paging, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(url)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	// Send out the HTTP request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(url)
		return nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(url)
		return nil, err
	}

	response := Response{}

	err = json.Unmarshal(b, &response)
	if err != nil {
		fmt.Println(url)
		return nil, err
	}

	err = json.Unmarshal(response.Data, &model)
	if err != nil {
		fmt.Println(url)
		return nil, err
	}

	return &response.Paging, nil
}
