/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"log"
	"os"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumble_ffmpeg"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/spf13/viper"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's current
// state.
type MumbleDJ struct {
	config         gumble.Config
	client         *gumble.Client
	keepAlive      chan bool
	defaultChannel []string
	audioStream    *gumble_ffmpeg.Stream
	logger         *log.Logger
	cache          *AudioCache
}

// OnConnect event. First moves MumbleDJ into the default channel specified
// via commandline args or configuration and moves to root channel if no default
// channel is provided.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	if dj.client.Channels.Find(dj.defaultChannel...) != nil {
		dj.client.Self.Move(dj.client.Channels.Find(dj.defaultChannel...))
	} else {
		dj.logger.Println("A default channel was not provided, moving to root channel.")
	}

	dj.audioStream = gumble_ffmpeg.New(dj.client)
	dj.audioStream.Volume = float32(viper.GetFloat64("DefaultVolume"))

	dj.client.Self.SetComment(viper.GetString("DefaultComment"))

	if viper.GetBool("CacheEnabled") {
		dj.cache.Update()
		go dj.cache.DeleteExpired()
	}
}

// OnDisconnect event. Terminates MumbleDJ thread if connection retry attempts
// fail.
func (dj *MumbleDJ) OnDisconnect(e *gumble.DisconnectEvent) {
	if e.Type == gumble.DisconnectError || e.Type == gumble.DisconnectKicked {
		dj.logger.Fatalln("Disconnected from server. Will retry connection in 30 second intervals for 15 minutes.")
		reconnectSuccess := false
		for retries := 0; retries <= 30; retries++ {
			dj.logger.Println("Retrying connection...")
			if err := dj.client.Connect(); err == nil {
				dj.logger.Println("Successfully reconnected to server.")
				reconnectSuccess = true
				break
			}
			time.Sleep(30 * time.Second)
		}
		if !reconnectSuccess {
			dj.logger.Fatalln("Could not reconnect to server. Exiting...")
			dj.keepAlive <- true
			os.Exit(1)
		}
	} else {
		dj.keepAlive <- true
	}
}

// OnTextMessage event. Checks for command prefix and calls ParseCommand if it exists.
// Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {
	plainMessage := gumbleutil.PlainText(&e.TextMessage)
	if len(plainMessage) != 0 {
		if plainMessage[0] == viper.GetString("CommandPrefix")[0] && plainMessage != viper.GetString("CommandPrefix") {
			// TODO: Call ParseCommand here once implemented.
			//ParseCommand(e.Sender, e.Sender.Name, plainMessage[1:])
		}
	}
}

// OnUserChange event. Checks UserChange type and adjusts skiplists to reflect the current
// status of the users of the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeDisconnected) {
		if dj.audioStream.IsPlaying() {
			// TODO: Remove playlist and clip skips here.
		}
	}
}

// SendPrivateMessage sends a private message to a user. This is a simple helper
// method that checks if the user is in the server before sending the private
// message.
func (dj *MumbleDJ) SendPrivateMessage(user *gumble.User, message string) {
	if targetUser := dj.client.Self.Channel.Users.Find(user.Name); targetUser != nil {
		targetUser.Send(message)
	}
}
