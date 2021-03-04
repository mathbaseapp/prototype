package service

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/mathml"
	"prototype.mathbase.app/repository"
	"prototype.mathbase.app/tokenizer"
)

// Tokenize db内のドキュメントに対してトークナイズする
func Tokenize() {

	parser := converter.GetParser(converter.Latex)
	tokenizer := &tokenizer.MathmlTokenizer{}
	processor := &QiitaArticleProcessor{Parser: parser, Tokenizer: tokenizer}

	repository.Documents.StreamEveryDocument(processor.Process)
}

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

// QiitaArticleProcessor qiitaの記事を処理する
type QiitaArticleProcessor struct {
	Parser    converter.Parser
	Tokenizer tokenizer.Tokenizer
}

var alphabet = regexp.MustCompile("^([a-z]|\\d+|\\+|\\-|=)$")

// Process 記事を処理する
func (q *QiitaArticleProcessor) Process(document repository.Document) error {
	if index, _ := repository.Indexes.SelectByID(document.ID); index != nil {
		return errors.New("すでにトークナイズが実行された記事です")
	}

	var indexes []*repository.Index
	formulas := q.drainFormula(document)
	for _, formula := range formulas {
		expr := formula.getValueInOneLine()
		res, err := q.Parser.Parse(expr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("以下のformulaのパースに失敗しました")
			fmt.Println(formula.getInfo())
			continue
		}

		for _, node := range res.Node.List() {
			base := 1.0
			if alphabet.MatchString(node.Value) {
				base = 0.0001
			}
			token := mathml.Printer(node)
			indexdoc := repository.IndexDocument{ID: document.ID, URL: document.URL, Title: document.Title}
			weight := 1.0 / float64(len(formulas)) * float64(utf8.RuneCountInString(token)) * base
			indexes = append(indexes, &repository.Index{Key: token, Location: strconv.Itoa(formula.startLine), Document: indexdoc, Weight: weight})
		}
	}

	_, err := repository.Indexes.InsertMany(indexes)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

var inlineReg = regexp.MustCompile(`([^\$]\$|^\$)([^\$]+)\$`) // $ ~ $ で囲まれる箇所
var displayReg = regexp.MustCompile(`\$\$([^\$]+)\$\$`)       // $$ ~ $$ で囲まれる箇所

func (q *QiitaArticleProcessor) drainFormula(document repository.Document) []formula {
	var tmp formula
	var formulas []formula

	body := document.Content
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
