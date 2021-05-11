package main

import (
	"context"
	"fmt"

	"github.com/barthr/newsapi"
)

func Get(cid string, m string) {
	hl, _ := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"google-news", "NPR"},
	})
	for _, a := range hl.Articles {
		ars = append(ars, a)
	}
	msg := fmt.Sprintf("Got %d articles", len(ars))
	Reply(cid, msg)
}

func Prev(cid string, s string) {
	if count == 0 {
		Reply(cid, "No more")
		return
	} else {
		count--
	}
	if len(ars) == 0 {
		Reply(cid, "Get articles first fool")
		return
	}
	Ereply(cid, AtoE(ars[count]))
	return
}

func Next(cid string, s string) {
	if count >= len(ars) {
		Reply(cid, "No more")
		return
	} else {
		count++
	}
	if len(ars) == 0 {
		Reply(cid, "Get articles first fool")
		return
	}
	Ereply(cid, AtoE(ars[count]))
}
