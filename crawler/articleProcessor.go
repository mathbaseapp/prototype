package crawler

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"prototype.mathbase.app/repository"

	"prototype.mathbase.app/tokenizer"

	"prototype.mathbase.app/converter"
)

// 一つのページの処理に責任を持つ
type articleProcessorInterface interface {
	drainFormula(article) []formula
	process(article) error
}

// QiitaArticleProcessor qiitaの記事を処理する
type QiitaArticleProcessor struct {
	Parser    converter.Parser
	Tokenizer tokenizer.Tokenizer
	util
}

// すでに保存されているページかどうかの確認など、複数のarticleProcessorで共通しそうな処理を持つ
type util struct {
}

// 保存済みの記事ならtrue まだならfalse
func (u *util) checkAlreadyStored(article article) bool {
	_, err := repository.Documents.SelectByURL(article.URL)
	return err == nil
}

func (q *QiitaArticleProcessor) process(article article) error {

	if q.checkAlreadyStored(article) {
		return errors.New("すでに保存されている記事です")
	}

	doc := &repository.Document{URL: article.URL, Title: article.Title, Content: article.Body}
	doc, err := repository.Documents.InsertOne(doc)
	if err != nil {
		fmt.Println(err)
		return errors.New("ページの保存に失敗しました")
	}

	formulas := q.drainFormula(article)
	for _, formula := range formulas {
		expr := formula.getValueInOneLine()
		res, err := q.Parser.Parse(expr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("以下のformulaのパースに失敗しました")
			fmt.Println(formula.getInfo())
			continue
		}

		tokens, err := q.Tokenizer.Tokenize(res.Node)
		if err != nil {
			fmt.Println(err)
			fmt.Println("formulaのtokenizeに失敗しました")
			continue
		}
		for _, token := range tokens {
			index := &repository.Index{Key: token, Location: strconv.Itoa(formula.startLine), Document: repository.IndexDocument{ID: doc.ID, URL: doc.URL, Title: doc.Title}}
			index, err = repository.Indexes.InsertOne(index)
			if err != nil {
				fmt.Println("index の保存時にエラーが発生しました")
			}
		}

	}
	return nil
}

var inlineReg = regexp.MustCompile(`([^\$]\$|^\$)([^\$]+)\$`) // $ ~ $ で囲まれる箇所
var displayReg = regexp.MustCompile(`\$\$([^\$]+)\$\$`)       // $$ ~ $$ で囲まれる箇所

func (q *QiitaArticleProcessor) drainFormula(article article) []formula {
	var tmp formula
	var formulas []formula

	body := article.Body
	mathFlg := false

	lines := strings.Split(body, "\n")
	for index, line := range lines {
		line = strings.TrimSpace(line)

		// ```math ~~ ``` で囲まれる部分を一行ずつのスライスにしてformulaに格納する
		if mathFlg && line == "```" {
			mathFlg = false
			formulas = append(formulas, tmp)
		}
		if mathFlg {
			tmp.lineLength++
			tmp.value = append(tmp.value, line)
		}
		if line == "```math" {
			tmp = *&formula{}
			tmp.startLine = index
			tmp.lineLength = 0
			mathFlg = true
		}

		// $ ~ $ で囲まれる部分
		if matches := inlineReg.FindAllStringSubmatch(line, -1); len(matches) > 1 {
			for _, match := range matches {
				formulas = append(formulas, *&formula{startLine: index, value: match[2:], lineLength: 1})
			}
		}

		// $$ ~ $$ で囲まれる部分
		if matches := displayReg.FindAllStringSubmatch(line, -1); len(matches) > 1 {
			for _, match := range matches {
				formulas = append(formulas, *&formula{startLine: index, value: match[1:], lineLength: 1})
			}
		}
	}

	return formulas
}
