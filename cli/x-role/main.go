package main

import (
	"github.com/cthulhu/x-role"
	"log"
)

func main() {
	if err := xrole.Run(); err != nil {
		log.Fatal(err)
	}
}
