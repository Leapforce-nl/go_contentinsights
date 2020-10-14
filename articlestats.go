package contentinsights

import (
	"fmt"
	"net/url"

	"cloud.google.com/go/civil"
)

// ArticleStat stores ArticleStat from contentinsights
//
type ArticleStat struct {
	ArticleID               int               `json:"article_id"`
	ArticleCreateDate       string            `json:"article_create_date"`
	ArticleURL              string            `json:"article_url"`
	ArticleTitle            string            `json:"article_title"`
	ArticleAuthors          *[]ArticleAuthor  `json:"article_authors"`
	ArticleSections         *[]ArticleSection `json:"article_sections"`
	ArticleTopics           *[]ArticleTopic   `json:"article_topics"`
	ArticleSubsType         string            `json:"article_subs_type"`
	ArticlePID              string            `json:"article_pid"`
	ArticleWordCount        int               `json:"article_word_count"`
	ArticleCreateTime       string            `json:"article_create_time"`
	ArticlesNumber          int               `json:"articles_number"`
	ArticleReads            int               `json:"article_reads"`
	AttentionMinutes        int               `json:"attention_minutes"`
	AttentionMinutesAverage int               `json:"attention_minutes_average"`
	SocialActions           int               `json:"social_actions"`
	ReadDepth               int               `json:"read_depth"`
	PageDepth               float64           `json:"page_depth"`
	ArticleReadsAverage     int               `json:"article_reads_average"`
	SocialActionsAverage    int               `json:"social_actions_average"`
	Visitors                int               `json:"visitors"`
	VisitorsNew             int               `json:"visitors_new"`
	VisitorsReturning       int               `json:"visitors_returning"`
	VisitorsLoyal           int               `json:"visitors_loyal"`
	RowsTotal               int               `json:"rows_total"`
	CPIGeneral              *int              `json:"cpi_general"`
	CPIExposure             *int              `json:"cpi_exposure"`
	CPIEngagement           *int              `json:"cpi_engagement"`
	CPILoyalty              *int              `json:"cpi_loyalty"`
	CPIExposureMarker       *string           `json:"cpi_exposure_marker"`
	CPIEngagementMarker     *string           `json:"cpi_engagement_marker"`
	CPILoyaltyMarker        *string           `json:"cpi_loyalty_marker"`
	CPIArticleReadsLimit    *int              `json:"cpi_article_reads_limit"`
	CPISampleSize           *int              `json:"cpi_sample_size"`
	CPIPerspective          *int              `json:"cpi_perspective"`
}

type ArticleAuthor struct {
	AuthorID int    `json:"author_id"`
	Name     string `json:"name"`
}

type ArticleSection struct {
	SectionID int    `json:"section_id"`
	Name      string `json:"name"`
}

type ArticleTopic struct {
	TopicID int    `json:"topic_id"`
	Name    string `json:"name"`
}

// GetArticleStats returns articleStats
//
func (ci *ContentInsights) GetArticleStats(dateFrom civil.Date, dateTo civil.Date, withChildren bool) (*[]ArticleStat, error) {
	values := url.Values{}
	values.Set("dimension", "article")
	values.Set("date_from", dateFrom.String())
	values.Set("date_to", dateTo.String())
	values.Set("limit", "1000")
	if withChildren {
		values.Set("with_children", "1")
	} else {
		values.Set("with_children", "0")

	}

	articleStats := []ArticleStat{}

	paging, err := ci.Get("stats", &values, &articleStats)
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
