package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	Id        int    `json:"id,omitempty"`
	Type      string `json:"type,omitempty"`
	Setup     string `json:"setup,omitempty"`
	Punchline string `json:"punchline,omitempty"`
}

func readJoke(body []byte) (*Joke, error) {
	var j = new(Joke)
	err := json.Unmarshal(body, &j)
	return j, err
}

func Jokes(cid string, m string, args []string) {
	url := "https://official-joke-api.appspot.com/random_joke"
	res, err := http.Get(url)
	if err != nil {
		Reply(cid, "Couldn't get a joke", nil)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Reply(cid, "Couldn't get a joke", nil)
		return
	}
	j, err := readJoke([]byte(body))
	Ereply(cid, JokeToEmbed(*j))
}
