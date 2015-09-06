/*
 * MumbleDJ
 * By Matthieu Grieger
 * configuration.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import "github.com/spf13/viper"

// LoadConfig loads the configuration file (if exists) and stores its values
// in Viper.
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.mumbledj")
	err := viper.ReadInConfig()
	if err != nil {

	}
}

// SetDefaultConfigValues sets the default configuration values for MumbleDJ
// using Viper.
func SetDefaultConfigValues() {
	// General
	viper.SetDefault("CommandPrefix", "!")
	viper.SetDefault("SkipRatio", 0.5)
	viper.SetDefault("PlaylistSkipRatio", 0.5)
	viper.SetDefault("DefaultComment", "Hello! I am a bot. Type !help for a list of commands.")
	viper.SetDefault("MaxSongDuration", 0)

	// Cache
	viper.SetDefault("CacheEnabled", false)
	viper.SetDefault("CacheMaximumSize", 512)
	viper.SetDefault("CacheExpiryTime", 24)

	// Volume
	viper.SetDefault("DefaultVolume", 0.2)
	viper.SetDefault("LowestVolume", 0.01)
	viper.SetDefault("HighestVolume", 0.8)

	// Aliases
	viper.SetDefault("AddAliases", []string{"add", "a"})
	viper.SetDefault("SkipAlises", []string{"skip", "s"})
	viper.SetDefault("SkipPlaylistAliases", []string{"skipplaylist", "sp"})
	viper.SetDefault("AdminSkipAlises", []string{"forceskip", "fs"})
	viper.SetDefault("AdminSkipPlaylistAliases", []string{"forceskipplaylist", "fsp"})
	viper.SetDefault("HelpAliases", []string{"help", "h"})
	viper.SetDefault("VolumeAliases", []string{"volume", "v"})
	viper.SetDefault("MoveAliases", []string{"move", "m"})
	viper.SetDefault("ReloadAliases", []string{"reload", "r"})
	viper.SetDefault("ResetAliases", []string{"reset", "re"})
	viper.SetDefault("NumClipsAliases", []string{"numclips", "numsongs", "nc", "ns"})
	viper.SetDefault("NextClipAliases", []string{"nextclip", "nextsong", "next"})
	viper.SetDefault("CurrentClipAliases", []string{"currentclip", "currentsong", "current"})
	viper.SetDefault("SetCommentAliases", []string{"setcomment", "comment"})
	viper.SetDefault("NumCachedAliases", []string{"numcached", "cached"})
	viper.SetDefault("CacheSizeAliases", []string{"cachesize", "size"})
	viper.SetDefault("KillAliases", []string{"kill", "k"})

	// Permissions
	viper.SetDefault("AdminsEnabled", true)
	viper.SetDefault("Admins", []string{"Matt"})
	viper.SetDefault("AddIsAdminCommand", false)
	viper.SetDefault("AddPlaylistIsAdminCommand", false)
	viper.SetDefault("SkipIsAdminCommand", false)
	viper.SetDefault("HelpIsAdminCommand", false)
	viper.SetDefault("VolumeIsAdminCommand", false)
	viper.SetDefault("MoveIsAdminCommand", true)
	viper.SetDefault("ReloadIsAdminCommand", true)
	viper.SetDefault("ResetIsAdminCommand", true)
	viper.SetDefault("NumClipsIsAdminCommand", false)
	viper.SetDefault("NextClipIsAdminCommand", false)
	viper.SetDefault("CurrentClipIsAdminCommand", false)
	viper.SetDefault("SetCommentIsAdminCommand", true)
	viper.SetDefault("NumCachedIsAdminCommand", true)
	viper.SetDefault("CacheSizeIsAdminCommand", true)
	viper.SetDefault("KillIsAdminCommand", true)
}
