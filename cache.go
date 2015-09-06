/*
 * MumbleDJ
 * By Matthieu Grieger
 * cache.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/viper"
)

// AudioCache is a struct that holds the number of audio clips currently
// cached and their combined file size.
type AudioCache struct {
	NumAudioClips int
	TotalFileSize int64
}

// NewAudioCache creates an empty AudioCache.
func NewAudioCache() *AudioCache {
	newCache := &AudioCache{
		NumAudioClips: 0,
		TotalFileSize: 0,
	}
	return newCache
}

// GetNumAudioClips returns the number of audio clips currently cached.
func (c *AudioCache) GetNumAudioClips() int {
	clips, _ := ioutil.ReadDir(fmt.Sprintf("%s/.mumbledj/audio", os.Getenv("HOME")))
	return len(clips)
}

// GetTotalFileSize calculates the total file size of the files within
// the cache and returns it.
func (c *AudioCache) GetTotalFileSize() int64 {
	var totalSize int64
	clips, _ := ioutil.ReadDir(fmt.Sprintf("%s/.mumbledj/audio", os.Getenv("HOME")))
	for _, clip := range clips {
		totalSize += clip.Size()
	}
	return totalSize
}

// CleanIfOverMaxSize checks the cache directory to determine if the filesize
// of the audio clips within exceed the user-specified size limit. If so, the
// oldest files get cleared until it is no longer exceeding the limit.
func (c *AudioCache) CleanIfOverMaxSize() {
	for c.GetTotalFileSize() > int64(viper.GetInt("CacheMaximumSize")*1048576) {
		if err := c.ClearOldest(); err != nil {
			break
		}
	}
}

// Update updates the AudioCache struct.
func (c *AudioCache) Update() {
	c.NumAudioClips = c.GetNumAudioClips()
	c.TotalFileSize = c.GetTotalFileSize()
}

// DeleteExpired deletes audio clips that are older than the cache period set
// within the user configuration.
func (c *AudioCache) DeleteExpired() {
	for range time.Tick(5 * time.Minute) {
		clips, _ := ioutil.ReadDir(fmt.Sprintf("%s/.mumbledj/audio", os.Getenv("HOME")))

	}
}
