package service

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"prototype.mathbase.app/crawler"
	"prototype.mathbase.app/repository"

	"prototype.mathbase.app/tokenizer"

	"prototype.mathbase.app/converter"
)

type formula struct {
	startLine  int // 何行目に現れたか
	lineLength int // 何行続いたか
	value      []string
}

func (f *formula) getInfo() string {
	var str string
	str += fmt.Sprintf("startLine: %d\t", f.startLine)
	str += fmt.Sprintf("lineLength: %d\n", f.lineLength)
	for _, v := range f.value {
		str += fmt.Sprintf("\t%s\n", v)
	}
	return str
}

func (f *formula) getValueInOneLine() string {
	var str string
	for _, v := range f.value {
		str += v + "\n"
	}
	return str
}

// 一つのページの処理に責任を持つ
type articleProcessorInterface interface {
	drainFormula(crawler.Article) []formula
	Process(crawler.Article) error
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
func (u *util) checkAlreadyStored(article crawler.Article) bool {
	_, err := repository.Documents.SelectByURL(article.URL)
	return err == nil
}

// Process 記事を処理する
func (q *QiitaArticleProcessor) Process(article crawler.Article) error {

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

func (q *QiitaArticleProcessor) drainFormula(article crawler.Article) []formula {
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
