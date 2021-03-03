package service

import (
	"fmt"
	"time"

	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/crawler"
	"prototype.mathbase.app/tokenizer"
)

// Crawle 記事をクロール・保存する
func Crawle() error {

	token := "04955c64db710699566b3420e4a8ae01ec907dd6"
	parser := converter.GetParser(converter.Latex)
	tokenizer := &tokenizer.MathmlTokenizer{}
	processor := &QiitaArticleProcessor{Parser: parser, Tokenizer: tokenizer}
	c := crawler.NewCrawler(token)
	for !c.Done {

		time.Sleep(time.Second * 5) // 1時間に1000回のアクセス制限に引っかからないよう止める
		articles, err := c.Crawle()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, article := range articles {
			err := processor.Process(article)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
