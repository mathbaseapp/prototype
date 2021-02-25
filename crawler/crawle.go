package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

func Crawle() {
	accessToken := "04955c64db710699566b3420e4a8ae01ec907dd6"
	per_page := "100" // 1ページあたりの記事数 1~100の間

	for _, tag := range TAGS {
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Second * 5) // 1時間に1000回のアクセス制限に引っかからないよう止める
			page := strconv.Itoa(i)

			url := "http://qiita.com/api/v2/tags/" + tag + "/items?per_page=" + per_page + "&page=" + page
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", "Bearer "+accessToken)

			client := new(http.Client)
			res, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			var articles []Article

			if err := json.Unmarshal(body, &articles); err != nil {
				log.Fatal(err)
			}

			for _, item := range articles {
				texs := getTex(item.Body)
				_ = texs
				fmt.Printf("\n\n%s に含まれるtexコードは以下の通り\n", item.Title)
				for i, tex := range texs {
					fmt.Printf("\nno.%d\n", i)
					for _, item := range tex {
						fmt.Printf("\t%s\n", item)
					}
				}
			}
		}
	}

}

var inlineReg = regexp.MustCompile(`[^\$]*\$([^\$]+)\$`) // $ ~ $ で囲まれる箇所
var displayReg = regexp.MustCompile(`\$\$([^\$]+)\$\$`)  // $$ ~ $$ で囲まれる箇所

func getTex(body string) [][]string {
	var texs [][]string
	var tex []string
	mathFlg := false

	lines := strings.Split(body, "\n")
	for _, line := range lines {
		// ```math ~~ ``` で囲まれる部分を一行ずつのスライスにしてtexに格納する
		if mathFlg && line == "```" {
			mathFlg = false
			texs = append(texs, tex)
			tex = make([]string, 0)
		}
		if mathFlg {
			tex = append(tex, line)
		}
		if line == "```math" {
			mathFlg = true
		}

		// $ ~ $ で囲まれる部分
		if matches := inlineReg.FindAllStringSubmatch(line, -1); len(matches) > 1 {
			for _, match := range matches {
				texs = append(texs, match[1:])
			}
		}

		// $$ ~ $$ で囲まれる部分
		if matches := displayReg.FindAllStringSubmatch(line, -1); len(matches) > 1 {
			for _, match := range matches {
				texs = append(texs, match[1:])
			}
		}
	}
	return texs
}
