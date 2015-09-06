/*
 * MumbleDJ
 * By Matthieu Grieger
 * service/youtube/clip.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jmoiron/jsonq"
)

// AudioClip holds the metadata for an audio clip extracted from a YouTube video.
type AudioClip struct {
	submitter string
	title     string
	id        string
	offset    int
	filename  string
	duration  string
	thumbnail string
	playlist  Playlist
}

// NewYouTubeAudioClip gathers the metadata for an audio clip extracted from a
// YouTube video and returns the song.
func NewYouTubeAudioClip(user, id, offset string, playlist *Playlist) (*AudioClip, error) {
	var apiResponse *jsonq.JsonQuery
	var err error
	var offsetDays, offsetHours, offsetMinutes, offsetSeconds int64
	var durationString string

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails&id=%s&key=%s",
		id, os.Getenv("YOUTUBE_API_KEY"))
	if apiResponse, err = PerformGetRequest(url); err != nil {
		return nil, err
	}

	title, _ := apiResponse.String("items", "0", "snippet", "title")
	thumbnail, _ := apiResponse.String("items", "0", "snippet", "thumbnails", "high", "url")
	duration, _ := apiResponse.String("items", "0", "contentDetails", "duration")

	if len(offset) != 0 {
		offsetDays, offsetHours, offsetMinutes, offsetSeconds = ParseOffset(offset)
	}
	totalOffset := int((offsetDays * 86400) + (offsetHours * 3600) + (offsetMinutes * 60) + offsetSeconds)

	days, hours, minutes, seconds := ParseDuration(duration)
	totalSeconds := int((days * 86400) + (hours * 3600) + (minutes * 60) + seconds)

	if hours != 0 {
		if days != 0 {
			durationString = fmt.Sprintf("%d:%02d:%02d:%02d", days, hours, minutes, seconds)
		} else {
			durationString = fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
		}
	} else {
		durationString = fmt.Sprintf("%d:%02d", minutes, seconds)
	}

	// TODO: Add check for max song duration configuration before returning.
	clip := &AudioClip{
		submitter: user,
		title:     title,
		id:        id,
		offset:    totalOffset,
		filename:  id + ".m4a",
		duration:  durationString,
		thumbnail: thumbnail,
		playlist:  nil,
	}

	return clip, nil
}

// Download downloads the audio clip via youtube-dl if it does not already exist on disk.
// All downloaded audio clips are stored in ~/.mumbledj/audio and should be automatically
// cleaned.
func (c *AudioClip) Download() error {
	// TODO: Implement Download().
}

// Delete deletes the audio clip from ~/.mumbledj/audio if the cache is disabled.
func (c *AudioClip) Delete() error {
	// TODO: Implement Delete().
}

// Submitter returns the name of the submitter of the AudioClip.
func (c *AudioClip) Submitter() string {
	return c.submitter
}

// Title returns the title of the AudioClip.
func (c *AudioClip) Title() string {
	return c.title
}

// ID returns the ID of the AudioClip.
func (c *AudioClip) ID() string {
	return c.id
}

// Filename returns the filename of the AudioClip.
func (c *AudioClip) Filename() string {
	return c.filename
}

// Duration returns the duration of the AudioClip.
func (c *AudioClip) Duration() string {
	return c.duration
}

// Thumbnail returns the thumbnail URL for the AudioClip.
func (c *AudioClip) Thumbnail() string {
	return c.thumbnail
}

// Playlist returns the playlist type for the AudioClip (may be nil).
func (c *AudioClip) Playlist() Playlist {
	return c.playlist
}

// ParseOffset parses a YouTube video offset and returns the parsed result in
// int64's of the offset days, hours, minutes, and seconds.
func ParseOffset(offset string) (int64, int64, int64, int64) {
	var offsetDays, offsetHours, offsetMinutes, offsetSeconds int64
	offsetRegex := regexp.MustCompile(`t\=(?P<days>\d+d)?(?P<hours>\d+h)?(?P<minutes>\d+m)?(?P<seconds>\d+s)?`)
	offsetMatch := offsetRegexp.FindStringSubmatch(offset)
	offsetResult := make(map[string]string)

	for i, name := range offsetRegexp.SubexpNames() {
		if i < len(offsetMatch) {
			offsetResult[name] = offsetMatch[i]
		}
	}

	if len(offsetResult["days"]) != 0 {
		offsetDays, _ = strconv.ParseInt(strings.TrimSuffix(offsetResult["days"], "d"), 10, 32)
	}
	if len(offsetResult["hours"]) != 0 {
		offsetHours, _ = strconv.ParseInt(strings.TrimSuffix(offsetResult["hours"], "h"), 10, 32)
	}
	if len(offsetResult["minutes"]) != 0 {
		offsetMinutes, _ = strconv.ParseInt(strings.TrimSuffix(offsetResult["minutes"], "m"), 10, 32)
	}
	if len(offsetResult["seconds"]) != 0 {
		offsetSeconds, _ = strconv.ParseInt(strings.TrimSuffix(offsetResult["seconds"], "s"), 10, 32)
	}

	return offsetDays, offsetHours, offsetMinutes, offsetSeconds
}

// ParseDuration parses the duration of a YouTube video and returns int64's of
// how long the video is in days, hours, minutes, and seconds.
func ParseDuration(duration string) (int64, int64, int64, int64) {
	var days, hours, minutes, seconds int64
	timestampRegexp := regexp.MustCompile(`P(?P<days>\d+D)?T(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	timestampMatch := timestampRegexp.FindStringSubmatch(duration)
	timestampResult := make(map[string]string)

	for i, name := range timestampRegexp.SubexpNames() {
		if i < len(timestampMatch) {
			timestampResult[name] = timestampMatch[i]
		}
	}

	if len(timestampResult["days"]) != 0 {
		days, _ = strconv.ParseInt(strings.TrimSuffix(timestampResult["days"], "D"), 10, 32)
	}
	if len(timestampResult["hours"]) != 0 {
		hours, _ = strconv.ParseInt(strings.TrimSuffix(timestampResult["hours"], "H"), 10, 32)
	}
	if len(timestampResult["minutes"]) != 0 {
		minutes, _ = strconv.ParseInt(strings.TrimSuffix(timestampResult["minutes"], "M"), 10, 32)
	}
	if len(timestampResult["seconds"]) != 0 {
		seconds, _ = strconv.ParseInt(strings.TrimSuffix(timestampResult["seconds"], "S"), 10, 32)
	}

	return days, hours, minutes, seconds
}
