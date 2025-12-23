package main

import (
	"fmt"
	"io"
	//"log"
	//"os"
)

var (
	r io.Reader
	w io.Writer
)

func init() {
	r, w = io.Pipe()
}

func main() {
	fmt.Println("Hello Main")

	initDiscordBot()
}

// creates the discord bot and a IO channel
func initDiscordBot() {
	//start the bot and give it the writer

	//check io for errors
}
