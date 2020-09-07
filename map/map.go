package main

import (
	"golang.org/x/tour/wc"
)

// Split 文章を単語ごとに区切り配列で返す。
func Split(article string) []string {
	var words []string
	var s string = ""
	for _, w := range article {
		if string(w) != " " {
			s += string(w)
		} else {
			words = append(words, s)
			s = ""
		}
	}
	words = append(words, s)
	return words
}

// WordCount 題意の関数。各単語が文章内にいくつずつ含まれているかをカウントしmapで返す。
func WordCount(article string) map[string]int {
	wordLists := Split(article)
	count := make(map[string]int)
	for _, w := range wordLists {
		count[w]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
