package service

import (
	"time"

	"prototype.mathbase.app/crawler"
	"prototype.mathbase.app/lg"
	"prototype.mathbase.app/repository"
)

// Crawle 記事をクロール・保存する
func Crawle() error {

	token := "04955c64db710699566b3420e4a8ae01ec907dd6"
	c := crawler.NewCrawler(token)
	for !c.Done {

		time.Sleep(time.Second * 5) // 1時間に1000回のアクセス制限に引っかからないよう止める
		articles, err := c.Crawle()
		if err != nil {
			lg.I.Println(err)
			continue
		}
		for _, article := range articles {
			if checkAlreadyStored(article) {
				continue
			}
			doc := &repository.Document{URL: article.URL, Title: article.Title, Content: article.Body}
			_, err = repository.Documents.InsertOne(doc)
			if err != nil {
				lg.I.Printf("%sの保存時にエラーが発生しました。\n", doc.URL)
				lg.I.Println(err)
			}
		}
	}
	return nil
}

// 保存済みの記事ならtrue まだならfalse
func checkAlreadyStored(article crawler.Article) bool {
	_, err := repository.Documents.SelectByURL(article.URL)
	return err == nil
}
