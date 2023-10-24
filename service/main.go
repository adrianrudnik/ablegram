package main

import (
	"git.jetbrains.space/the/ablegram/apiserver"
	"git.jetbrains.space/the/ablegram/parser"
)

func main() {
	_, err := parser.ParseAls(".samples/sample-001-v11-empty.als")
	//_, err := parser.ParseAls(".samples/800-ios-note-casolare.als")
	if err != nil {
		panic(err)
	}

	apiserver.Serve(":10000")
}
