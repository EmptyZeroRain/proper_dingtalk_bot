package main

import (
	"log"
	"os"

	"github.com/EmptyZeroRain/proper_dingtalk_bot/cmd/dingtalk"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	dingtalk.Execute()
}
