package main

import (
	"github.com/kyokomi/emoji/v2"
)

func main(){
GetMessage()
}


func GetMessage() string {
	return emoji.Sprintf("Hello 🗺️!")
}
