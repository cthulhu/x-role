package main

import (
	"github.com/cthulhu/x-role"
	"log"
)

func main(){
	if err := xrole.Assume(); err != nil {
		log.Fatal(err)
	}
}