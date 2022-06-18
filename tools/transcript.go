package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func ParseTranscript(id string) {
	handleItems("../transcripts/" + id + ".html")
}

func handleItems(filename string) {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	tkn := html.NewTokenizer(strings.NewReader(s))

	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:

			t := tkn.Token()
			if t.Data == "tr" {
			} else if t.Data == "td" {
			}

		case tt == html.TextToken:

			t := tkn.Token()
			txt := strings.TrimSpace(t.Data)
			if len(txt) == 0 {
				continue
			}
			if !unicode.IsLetter(rune(txt[0])) {
				break
			}
			fmt.Println(txt)
			if txt == "" {
				continue
			}

		}

	}

	return
}
