/*
 * MumbleDJ
 * By Matthieu Grieger
 * service/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package service

// Playlist interface. Each service will implement these methods in their
// Playlist types.
type Playlist interface {
	ID() string
	Title() string
	Service() string
}
