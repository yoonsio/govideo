package models

import (
	"sync"
	"time"
)

// Media -
type Media struct {
	Path   string `bson:"_id"`
	Name   string
	Size   int64
	Access []string
	Added  time.Time
}

var mediaPool = sync.Pool{
	New: func() interface{} {
		return &Media{}
	},
}

// GetMedia gets Media struct from sync pool
func GetMedia() *Media {
	return mediaPool.Get().(*Media)
}

// RecycleMedia puts back Media struct back into sync pool
func RecycleMedia(media *Media) {
	mediaPool.Put(media)
}

// MediaList -
//easyjson:json
type MediaList struct {
	Data  []Media
	Count int
}

var mediaListPool = sync.Pool{
	New: func() interface{} {
		return &MediaList{}
	},
}

// GetMediaList gets MediaList struct from sync pool
func GetMediaList() *MediaList {
	return mediaListPool.Get().(*MediaList)
}

// RecycleMediaList puts back MediaList struct back into sync pool
func RecycleMediaList(mediaList *MediaList) {
	mediaListPool.Put(mediaList)
}
