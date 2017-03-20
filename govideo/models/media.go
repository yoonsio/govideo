package models

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Media -
type Media struct {
	Path   string `bson:"_id"`
	Name   string
	Size   int64
	Access []bson.ObjectId
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
