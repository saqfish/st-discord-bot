package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/barthr/newsapi"
)

var nc *newsapi.Client
var articles []newsapi.Article
var count int
var source string

func init() {
	nc = newsapi.NewClient(os.Args[2], newsapi.WithHTTPClient(http.DefaultClient))
}

func News(cid string, args ...string) {
	articles = nil

	hl, err := nc.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{source},
	})

	if err != nil || len(args) < 1 {
		Reply(cid, "Couldn't get news")
		return
	}

	for _, a := range hl.Articles {
		articles = append(articles, a)
	}

	msg := fmt.Sprintf("Got %d articles, use article to select article or next & prev to cycle.", len(articles))
	Reply(cid, msg)
}

func Article(cid string, args ...string) {
	n, err := strconv.ParseInt(args[0], 10, 62)
	num := int(n)
	if err != nil || len(args) < 1 {
		Reply(cid, "Invalid arg")
		return
	}
	if num > len(articles) || num == 0 {
		msg := fmt.Sprintf("Out of range, there are %d articles", len(articles))
		Reply(cid, msg)
		return
	}
	if len(articles) < 1 {
		Reply(cid, "No articles available")
		return
	}
	if int(num) > len(articles) || int(num) < 1 {
		Reply(cid, "Invalid arg")
		return
	}
	count = num
	Ereply(cid, ArticleToEmbed(articles[num-1]))
}

func Prev(cid string, args ...string) {
	if count == 0 {
		Reply(cid, "No more")
		return
	} else {
		count--
	}
	if len(articles) == 0 {
		Reply(cid, "Get articles first fool")
		return
	}
	Ereply(cid, ArticleToEmbed(articles[count]))
	return
}

func Next(cid string, args ...string) {
	if count >= len(articles) {
		Reply(cid, "No more")
		return
	} else {
		count++
	}
	if len(articles) == 0 {
		Reply(cid, "Get articles first fool")
		return
	}
	Ereply(cid, ArticleToEmbed(articles[count]))
}

func Source(cid string, args ...string) {
	if len(args) < 1 {
		Reply(cid, "Pick a source: cbs-news, fox-news, google-news, nbc-news, nfl-news, cnn, engadget")
		return
	}
	source = args[0]
	msg := fmt.Sprintf("Source set to %s", args[0])
	Reply(cid, msg)
}
