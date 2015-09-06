/*
 * MumbleDJ
 * By Matthieu Grieger
 * service/clip.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package service

// AudioClip interface. Each service will implement these methods in their
// AudioClip types.
type AudioClip interface {
	Download() error
	Delete() error
	Submitter() string
	Title() string
	ID() string
	Filename() string
	Duration() string
	Thumbnail() string
	Playlist() Playlist
}
