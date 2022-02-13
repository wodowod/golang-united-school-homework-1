package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	msg := "Hello 🗺️!"
	return emoji.Sprintf(msg)
}
