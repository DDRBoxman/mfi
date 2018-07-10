package main

import (
	"github.com/DDRBoxman/mfi"
	"log"
)

func main() {
	client, err := mfi.MakeMFIClient("10.42.42.12")
	if err != nil {
		log.Panic(err)
	}

	err = client.Auth("ubnt", "ubnt")
	if err != nil {
		log.Panic(err)
	}

	err = client.SetOutputEnabled(5, false)
	if err != nil {
		log.Panic(err)
	}
}