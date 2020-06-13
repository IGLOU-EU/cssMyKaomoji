package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var cssFile = "kaomoji.css"
var htmlFile = "test/index.html"
var kaomoji map[string]string

func main() {
	start := time.Now()
	kaomoji = make(map[string]string)

	res, err := http.Get("http://kaomoji.ru/en/")
	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	extractKaomoji(string(page))
	makeCSSFile(kaomoji)
	makeDemoFile(kaomoji)

	fmt.Printf("In %s\n%d Kaomiji generated\nCSS file: ./%s\nPreview file: ./%s\n", time.Since(start), len(kaomoji), cssFile, htmlFile)
}

func extractKaomoji(p string) {
	i, bf, cat := 0, "", ""
	sc := bufio.NewScanner(strings.NewReader(p))

	reCat := regexp.MustCompile(`(?m)>([^<>]+)<`)
	reSpe := regexp.MustCompile(`(?m) |\.|\/|'`)

	_ = i
	for sc.Scan() {
		l := sc.Text()

		if strings.Contains(l, "<h3><a name=") {
			cat = reCat.FindStringSubmatch(l)[1]

			if cat == "Special" {
				i = -1
			} else {
				i = 1
			}

			continue
		}

		if strings.Contains(l, "<td><span>") {
			bf = strings.Replace(strings.TrimSpace(l), "<td><span>", "", 1)
			bf = strings.Replace(bf, "</span></td>", "", 1)
			bf = strings.ReplaceAll(bf, "\\", "\\\\")
			bf = strings.ReplaceAll(bf, "\"", "\\\"")

			if bf == "" || i == -1 {
				continue
			}

			kaomoji[cat+"_"+strconv.Itoa(i)] = bf
			i++
		}

		if strings.Contains(l, "<td style=\"font-family:Comic Sans MS\">") {
			cat = strings.Replace(strings.TrimSpace(l), "<td style=\"font-family:Comic Sans MS\">", "", 1)
			cat = strings.Replace(cat, "</td>", "", 1)
			cat = reSpe.ReplaceAllString(cat, "_")

			if bf == "" {
				continue
			}

			kaomoji[cat] = bf
		}
	}
}

func makeCSSFile(kl map[string]string) {
	css := `.kaomoji {font-family: "Monaco", "DejaVu Sans Mono", "Lucida Console", "Andale Mono", "monospace";}`

	f, err := os.OpenFile(cssFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for k, i := range kl {
		css += ".kaomoji." + k + "::before{content:\"" + i + "\"}"
	}

	if _, err = f.WriteString(css); err != nil {
		log.Fatal(err)
	}
}

func makeDemoFile(kl map[string]string) {
	css := "<!DOCTYPE html><html><head><meta charset=\"utf-8\"><link rel='stylesheet' type='text/css' href='../kaomoji.css'><style>body{display:grid;grid-template-columns:repeat(3,1fr);grid-gap:.5rem;}span{padding:1rem;border:2px dashed lightblue;text-align:center; min-height: 8em;display: flex;flex-direction: column;justify-content: space-around;white-space: nowrap;}sup{padding-top: .5em;margin-top: .5em;border-top: 1px dashed lightblue;display:block}::before{font-weight: bold;font-size: 1.5em;}</style></head><body>"

	f, err := os.OpenFile(htmlFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for k := range kl {
		css += "<span role=\"image\" class=\"kaomoji " + k + "\"><sup>." + k + "</sup></span>"
	}

	css += "</body></html>"

	if _, err = f.WriteString(css); err != nil {
		log.Fatal(err)
	}
}
