package main

var Commands = map[string]func(cid string, args ...string){
	"repeat":  Repeat,
	"say":     Repeat,
	"xkcd":    Xkcd,
	"source":  Source,
	"news":    News,
	"article": Article,
	"next":    Next,
	"prev":    Prev,
	"joke":    Jokes,
}
