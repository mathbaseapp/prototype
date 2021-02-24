package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

func Crawle() {
	tag := "math"
	res, err := http.Get("http://qiita.com/api/v2/tags/" + tag + "/items")
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
		fmt.Printf("\n\n%s に含まれるtexコードは以下の通り\n", item.Title)
		for i, tex := range texs {
			fmt.Printf("\nno.%d\n", i)
			for _, item := range tex {
				fmt.Printf("\t%s\n", item)
			}
		}
		// fmt.Printf("%s\n", item.Body)
	}

}

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
	}
	return texs
}
