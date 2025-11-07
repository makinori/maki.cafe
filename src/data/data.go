package data

import "github.com/makinori/foxlib/foxcache"

func Init() {
	foxcache.Init("cache", []foxcache.DataInterface{
		&Anilist, &Squirrels,
	})
}
