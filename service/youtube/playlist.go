/*
 * MumbleDJ
 * By Matthieu Grieger
 * service/youtube/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import (
	"fmt"
	"os"

	"github.com/jmoiron/jsonq"
)

// Playlist holds the metadata for a YouTube playlist.
type Playlist struct {
	id      string
	title   string
	service string
}

// NewYouTubePlaylist gathers the metadata for a YouTube playlist and returns it.
func NewYouTubePlaylist(user, id string) (*Playlist, error) {
	var apiResponse *jsonq.JsonQuery
	var err error

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?part=snippet&id=%s&key=%s",
		id, os.Getenv("YOUTUBE_API_KEY"))
	if apiResponse, err = PerformGetRequest(url); err != nil {
		return nil, err
	}
}
