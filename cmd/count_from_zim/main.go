package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	zim "github.com/akhenakh/gozim"
	"github.com/microcosm-cc/bluemonday"
)

const NEEDED_MIME_TYPE string = "text/html"

var z *zim.ZimReader
var myString string
var text string
var counts = make(map[string]int)

func printHelp() {
	fmt.Println("Usage: command file")
	os.Exit(0)
}

func normalizeData(src_string string) string {
	policy := bluemonday.StripTagsPolicy()
	result := policy.Sanitize(src_string)
	result = strings.Join(strings.Fields(result), " ")
	return result
}
func isNeededArticle(a *zim.Article) bool {
	if a.EntryType == zim.DeletedEntry {
		return false
	}
	if a.Namespace != 'A' {
		return false
	}
	if a.Title == "" {
		return false
	}
	if a.MimeType() != NEEDED_MIME_TYPE {
		return false
	}

	return true
}

func parseData(src string) map[string]int {
	result := make(map[string]int)
	for _, item := range strings.Fields(src) {
		_, has := result[item]
		if has {
			result[item]++
		} else {
			result[item] = 1
		}
	}
	return result
}

func updateMap(src map[string]int, update map[string]int) map[string]int {
	for k, v := range update {
		_, has := src[k]
		if has {
			src[k] = src[k] + v
		} else {
			src[k] = v
		}
	}
	return src
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}
	filePath := os.Args[1]

	z, err := zim.NewReader(filePath, false)
	if err != nil {
		panic(err)
	}
	i := 0
	z.ListTitlesPtrIterator(func(idx uint32) {
		a, err := z.ArticleAtURLIdx(idx)
		if err != nil {
			log.Fatal(err)
		}
		if !isNeededArticle(a) {
			i++
			return
		}
		data, err := a.Data()
		if err != nil {
			log.Fatal(err.Error())
		}
		text = normalizeData(string(data[:]))
		counts = updateMap(counts, parseData(text))
	})
	for k, v := range counts {
		fmt.Println(k, " ", v)
	}

}
