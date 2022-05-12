package main

import (
	"github.com/kyokomi/emoji/v2"
)

func Hello() string {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	return helloMessage
}
