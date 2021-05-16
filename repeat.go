package main

func Repeat(cid string, args ...string) {
	session.ChannelMessageSend(cid, args[0])
}
