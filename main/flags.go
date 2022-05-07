package main

import "flag"

var botToken, webhookUrl, listenAddr string

func init() {
	flag.StringVar(&botToken, "token", "", "tg bot token")
	flag.StringVar(&webhookUrl, "webhook", "", "tg bot webhook url")
	flag.StringVar(&listenAddr, "listen", "127.0.0.1:3456", "listen address")
}
