/*
 * MumbleDJ
 * By Matthieu Grieger
 * queue.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"errors"
	"fmt"

	"github.com/matthieugrieger/mumbledj/service"
	"github.com/spf13/viper"
)

// AudioQueue is a type that holds a slice containing the audio clips that make
// up the current queue.
type AudioQueue struct {
	queue []service.AudioClip
}

// NewAudioQueue initializes a new queue and returns it.
func NewAudioQueue() *AudioQueue {
	return &AudioQueue{
		queue: make([]service.AudioClip, 0),
	}
}

// AddAudioClip adds an AudioClip to the AudioQueue.
func (q *AudioQueue) AddAudioClip(a service.AudioClip) error {
	beforeLen := len(q.queue)
	q.queue = append(q.queue, a)
	if len(q.queue) == beforeLen+1 {
		dj.logger.Println(fmt.Sprintf("Added a %s clip named %s to the queue.", a.Service(), a.Title()))
		return nil
	}
	return errors.New("Could not add audio clip to queue.")
}

// CurrentAudioClip returns the current AudioClip if exists.
func (q *AudioQueue) CurrentAudioClip() service.AudioClip {
	if len(q.queue) != 0 {
		return q.queue[0]
	}
	return nil
}

// NextAudioClip moves to the next AudioClip in the AudioQueue and removes the
// first AudioClip in the queue.
func (q *AudioQueue) NextAudioClip() {
	currentClip := q.CurrentAudioClip()
	if currentClip == nil {
		return
	}
	if currentClip.Playlist() != nil {
		if a, err := q.PeekNext(); err == nil {
			if currentClip.Playlist().ID() != a.Playlist().ID() {
				ResetPlaylistSkips()
			}
		} else {
			ResetPlaylistSkips()
		}
	}
	q.queue = q.queue[1:]
}

// PeekNext peeks at the next AudioClip and returns it.
func (q *AudioQueue) PeekNext() (service.AudioClip, error) {
	if len(q.queue) > 1 {
		return q.queue[1], nil
	}
	return nil, errors.New("There isn't an audio clip coming up next.")
}

// Traverse is a traversal function for AudioQueue. Allows a visit function
// to be passed in which performs the specified action on each queue item.
func (q *AudioQueue) Traverse(visit func(i int, a service.AudioClip)) {
	for aQueue, queueAudioClip := range q.queue {
		visit(aQueue, queueAudioClip)
	}
}

// OnAudioClipFinished event. Deletes AudioClip that just finished playing if
// caching is disabled and then queues the next AudioClip if one exists.
func (q *AudioQueue) OnAudioClipFinished(wasPlaylistSkip bool) {
	currentClip := q.CurrentAudioClip()
	var playlistID string
	if currentClip.Playlist() != nil {
		playlistID = currentClip.Playlist().ID()
	} else {
		playlistID = ""
	}

	if !viper.GetBool("CacheEnabled") {
		q.CurrentAudioClip().Delete()
	}
	q.NextAudioClip()

	done := false
	for len(q.queue) != 0 && !done {
		newAudioClip := q.CurrentAudioClip()
		if wasPlaylistSkip && len(playlistID) != 0 && newAudioClip.Playlist() != nil {
			if newAudioClip.Playlist().ID() == currentClip.Playlist().ID() {
				q.NextAudioClip()
			} else {
				done = true
			}
		} else {
			done = true
		}
	}

	if len(q.queue) != 0 {
		if err := q.CurrentAudioClip().Download(); err == nil {
			// TODO: Play new audio clip here.
		} else {
			// TODO: Log error message as download has failed. Also notify
			// user that the download for their audio clip has failed.
			q.OnAudioClipFinished(false)
		}
	}
}
