package main

import (
	"fmt"
	"log"

	"gitlab.com/idoko/foxtop/db"
)

func main() {

	dbFile := "tmp.sqlite"
	db, err := db.Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	store, err := db.AllHosts()
	if err != nil {
		log.Fatal(err)
	}

	for _, host := range store.Hosts() {
		fmt.Println(host.HostName())
	}
}
