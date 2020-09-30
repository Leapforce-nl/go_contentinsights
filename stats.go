package contentinsights

import (
	"fmt"
	"net/url"

	"cloud.google.com/go/civil"
)

// Stat stores Stat from contentinsights
//
type Stat struct {
	ArticleID int `json:"article_id"`
}

// GetStats returns stats
//
func (ci *ContentInsights) GetStats(dimension string, dateFrom civil.Date, dateTo civil.Date, withChildren bool) (*[]Stat, error) {
	values := url.Values{}
	values.Set("dimension", dimension)
	values.Set("date_from", dateFrom.String())
	values.Set("date_to", dateTo.String())
	if withChildren {
		values.Set("with_children", "1")
	} else {
		values.Set("with_children", "0")

	}

	stats := []Stat{}

	paging, err := ci.Get("stats", &values, &stats)
	if err != nil {
		fmt.Println("ERROR in GetStats:", err)
		return nil, err
	}

	for paging.Next != nil {
		_stats := []Stat{}

		paging, err = ci.GetURL(*paging.Next, &_stats)
		if err != nil {
			fmt.Println("ERROR in GetStats:", err)
			fmt.Println("url:", *paging.Next)
			return nil, err
		}

		stats = append(stats, _stats...)
	}

	return &stats, nil
}
