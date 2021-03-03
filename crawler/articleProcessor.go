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
}

func (q *QiitaArticleProcessor) process(article article) error {

	doc := &repository.Document{URL: article.URL, Title: article.Title, Content: "hoge"}
	doc, err := repository.Documents.InsertOne(doc)
	if err != nil {
		fmt.Println(err)
		return errors.New("ページの保存に失敗しました")
	}

	formulas := q.drainFormula(article)
	for _, formula := range formulas {
		if formula.lineLength == 1 {
			res, err := q.Parser.Parse(formula.value[0])
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
				if token == "<></>" {
					//TODO: parse した後にres.Nodeに何も入ってない場合に起こる。parse側でerrを返してcontinueするのがベター
					continue
				}
				index := &repository.Index{Key: token, Location: strconv.Itoa(formula.startLine), Document: repository.IndexDocument{ID: doc.ID, URL: doc.URL, Title: doc.Title}}
				index, err = repository.Indexes.InsertOne(index)
				if err != nil {
					fmt.Println("index の保存時にエラーが発生しました")
				}
				fmt.Println(token)
			}
			fmt.Println("")
			// fmt.Println(mathml.Printer(res.Node))
		}
		// fmt.Println(formula.getInfo())
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
