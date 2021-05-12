package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/barthr/newsapi"
)

func News(cid string, m string, args []string) {
	hl, _ := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"google-news", "NPR"},
	})
	ars = nil
	for _, a := range hl.Articles {
		ars = append(ars, a)
	}
	msg := fmt.Sprintf("Got %d articles, use article to select article or next & prev to cycle.", len(ars))
	Reply(cid, msg, nil)
}

func Article(cid string, s string, args []string) {
	n, err := strconv.ParseInt(args[0], 10, 62)
	num := int(n)
	if err != nil || len(args) < 1 {
		Reply(cid, "Invalid arg", nil)
		return
	}
	if num > len(ars) || num == 0 {
		msg := fmt.Sprintf("Out of range, there are %d articles", len(ars))
		Reply(cid, msg, nil)
		return
	}
	if len(ars) < 1 {
		Reply(cid, "No articles available", nil)
		return
	}
	if int(num) > len(ars) || int(num) < 1 {
		Reply(cid, "Invalid arg", nil)
		return
	}
	count = num
	Ereply(cid, AtoE(ars[num-1]))
}

func Prev(cid string, s string, args []string) {
	if count == 0 {
		Reply(cid, "No more", nil)
		return
	} else {
		count--
	}
	if len(ars) == 0 {
		Reply(cid, "Get articles first fool", nil)
		return
	}
	Ereply(cid, AtoE(ars[count]))
	return
}

func Next(cid string, s string, args []string) {
	if count >= len(ars) {
		Reply(cid, "No more", nil)
		return
	} else {
		count++
	}
	if len(ars) == 0 {
		Reply(cid, "Get articles first fool", nil)
		return
	}
	Ereply(cid, AtoE(ars[count]))
}
