/*
 * MumbleDJ
 * By Matthieu Grieger
 * main.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var dj = MumbleDJ{
	keepAlive: make(chan bool),
	queue:     NewAudioQueue(),
	cache:     NewAudioCache(),
	logger:    log.New(os.Stdout, "[MumbleDJ] ", log.Lshortfile),
}

func main() {
	if !viper.IsSet("YouTubeAPIKey") {
		dj.logger.Fatalln("You do not have a YouTube API key defined in your configuration.")
	}
}
