package crawler

import (
	"fmt"
	"regexp"
	"strings"
)

// 一つのページの処理に責任を持つ
type articleProcessorInterface interface {
	drainFormula(article) []formula
	process(article) error
}

// QiitaArticleProcessor qiitaの記事を処理する
type QiitaArticleProcessor struct {
}

func (q *QiitaArticleProcessor) process(article article) error {
	formulas := q.drainFormula(article)
	for _, formula := range formulas {
		fmt.Println(formula.getInfo())
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
