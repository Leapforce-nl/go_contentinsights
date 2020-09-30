package contentinsights

import (
	"fmt"
	"net/url"

	"cloud.google.com/go/civil"
)

// ArticleStat stores ArticleStat from contentinsights
//
type ArticleStat struct {
	ArticleID int `json:"article_id"`
}

// GetArticleStats returns articleStats
//
func (ci *ContentInsights) GetArticleStats(dateFrom civil.Date, dateTo civil.Date, withChildren bool) (*[]ArticleStat, error) {
	values := url.Values{}
	values.Set("dimension", "article")
	values.Set("date_from", dateFrom.String())
	values.Set("date_to", dateTo.String())
	if withChildren {
		values.Set("with_children", "1")
	} else {
		values.Set("with_children", "0")

	}

	articleStats := []ArticleStat{}

	paging, err := ci.Get("articleStats", &values, &articleStats)
	if err != nil {
		fmt.Println("ERROR in GetArticleStats:", err)
		return nil, err
	}

	for paging.Next != nil {
		_articleStats := []ArticleStat{}

		paging, err = ci.GetURL(*paging.Next, &_articleStats)
		if err != nil {
			fmt.Println("ERROR in GetArticleStats:", err)
			fmt.Println("url:", *paging.Next)
			return nil, err
		}

		articleStats = append(articleStats, _articleStats...)
	}

	return &articleStats, nil
}
