package main

import (
	"git.jetbrains.space/the/ablegram/parser"
)

func main() {
	_, err := parser.ParseAls(".samples/800-ios-note-casolare.als")
	if err != nil {
		panic(err)
	}
}
