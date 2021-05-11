package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Comic struct {
	Month      string `json:"month,omitempty"`
	Num        int    `json:"num,omitempty"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year,omitempty"`
	News       string `json:"news,omitempty"`
	Safe_title string `json:"safe_title,omitempty"`
	Transcript string `json:"transcript,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Img        string `json:"img,omitempty"`
	Title      string `json:"title,omitempty"`
	Day        string `json:"day,omitempty"`
}

func readComic(body []byte) (*Comic, error) {
	var s = new(Comic)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Sorry, It didn't work", err)
	}
	return s, err
}
func Xkcd(cid string, m string) {
	url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", rand.Intn(600))
	res, err := http.Get(url)
	if err != nil {
		Reply(cid, "Couldn't get comic")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Reply(cid, "Couldn't get comic")
	}
	r, err := readComic([]byte(body))
	Ereply(cid, CtoE(*r))
}
