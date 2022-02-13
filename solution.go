package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	msg := "🗺️"
	fullmsg := emoji.Sprint("Hello %v!", msg)
	return fullmsg
}
