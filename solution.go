package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	return helloMessage
}
