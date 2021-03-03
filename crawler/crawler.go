package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Crawler クローラーの実体
type Crawler struct {
	token  string
	urlMap map[string]bool

	currentTagIndex int
	finalTagIndex   int

	currentPageIndex int
	finalPageIndex   int

	Done bool
}

// NewCrawler コンストラクタ
func NewCrawler(token string) *Crawler {
	c := new(Crawler)

	c.token = token
	c.currentTagIndex = 0
	c.finalTagIndex = len(TAGS) - 1

	c.currentPageIndex = 1
	c.finalPageIndex = 100
	c.urlMap = make(map[string]bool)
	c.Done = false

	return c
}

// Crawle クロール
func (c *Crawler) Crawle() ([]Article, error) {
	articles, err := c.getArticles()

	if err != nil {
		fmt.Println(err)
		fmt.Printf("currentTag: %s\tcurrentTagIndex: %d\tcurrentPageIndex: %d\n", TAGS[c.currentTagIndex], c.currentTagIndex, c.currentPageIndex)
		c.goNextPage()
		return nil, err
	}

	if len(articles) == 0 {
		c.goNextTag()
		return nil, err
	}

	for _, article := range articles {
		if c.urlMap[article.URL] {
			continue
		}
		c.urlMap[article.URL] = true
	}

	c.goNextPage()
	return articles, nil
}

func (c *Crawler) getArticles() ([]Article, error) {
	url := "http://qiita.com/api/v2/tags/" + TAGS[c.currentTagIndex] + "/items?per_page=" + strconv.Itoa(c.finalPageIndex) + "&page=" + strconv.Itoa(c.currentPageIndex)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("レスポンスの取得に失敗")
		return make([]Article, 0), err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("レスポンスの読み込みに失敗")
		return make([]Article, 0), err
	}

	var articles []Article

	if err := json.Unmarshal(body, &articles); err != nil {
		fmt.Println("レスポンスのパースに失敗")
		fmt.Println(string(body))
		return make([]Article, 0), err
	}

	return articles, nil
}

func (c *Crawler) goNextTag() {
	c.currentTagIndex++
	if c.currentTagIndex > c.finalTagIndex {
		c.Done = true
	}
}

func (c *Crawler) goNextPage() {
	c.currentPageIndex++
	if c.currentPageIndex > c.finalPageIndex {
		c.goNextTag()
		c.currentPageIndex = 1
	}
}
