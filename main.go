package main

import (
	"github.com/ikun666/go_webserver/cmd"
)

func main() {
	cmd.Start()
	defer cmd.End()
}
