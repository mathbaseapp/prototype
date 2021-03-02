package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Crawler クローラーの実体
type Crawler struct {
	token  string
	urlMap map[string]bool

	currentTagIndex int
	finalTagIndex   int

	currentPageIndex int
	finalPageIndex   int

	done bool

	articleProcessor articleProcessorInterface
}

// NewCrawler コンストラクタ
func NewCrawler(token string, ap articleProcessorInterface) *Crawler {
	c := new(Crawler)

	c.token = token
	c.articleProcessor = ap
	c.currentTagIndex = 0
	c.finalTagIndex = len(TAGS) - 1

	c.currentPageIndex = 1
	c.finalPageIndex = 100
	c.urlMap = make(map[string]bool)
	c.done = false

	return c
}

// Crawle クロール
func (c *Crawler) Crawle() {
	for !c.done {
		time.Sleep(time.Second * 5) // 1時間に1000回のアクセス制限に引っかからないよう止める
		articles, err := c.getArticles()

		if err != nil {
			fmt.Println(err)
			fmt.Printf("現在のCrawlerの状態は以下の通りです\n%#v", c)
			c.goNextPage()
			continue
		}

		if len(articles) == 0 {
			c.goNextTag()
			continue
		}

		for _, article := range articles {
			if c.urlMap[article.URL] {
				continue
			}
			c.urlMap[article.URL] = true
			err = c.articleProcessor.process(article)
			if err != nil {
				fmt.Printf("%s の処理に失敗しました\n", article.Title)
				fmt.Println(err)
			}
		}

		c.goNextPage()
	}
}

func (c *Crawler) getArticles() ([]article, error) {
	url := "http://qiita.com/api/v2/tags/" + TAGS[c.currentTagIndex] + "/items?per_page=" + strconv.Itoa(c.finalPageIndex) + "&page=" + strconv.Itoa(c.currentPageIndex)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.token)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("レスポンスの取得に失敗")
		return make([]article, 0), err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("レスポンスの読み込みに失敗")
		return make([]article, 0), err
	}

	var articles []article

	if err := json.Unmarshal(body, &articles); err != nil {
		fmt.Println("レスポンスのパースに失敗")
		fmt.Println(string(body))
		return make([]article, 0), err
	}

	return articles, nil
}

func (c *Crawler) goNextTag() {
	c.currentTagIndex++
	if c.currentTagIndex > c.finalTagIndex {
		c.done = true
	}
}

func (c *Crawler) goNextPage() {
	c.currentPageIndex++
	if c.currentPageIndex > c.finalPageIndex {
		c.goNextTag()
		c.currentPageIndex = 1
	}
}
