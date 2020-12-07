package main

import (
	"fmt"
	"gitlab.com/idoko/foxtop/api"
	"gitlab.com/idoko/foxtop/db"
	"log"
	"os"
)

var ListenAddr = "localhost:8080"

func main() {
	profileDir := "."
	dbFile, err := getDbFile(profileDir)
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	router := api.InitRouter(db)
	router.Run(ListenAddr)
}

func getDbFile(profileDir string) (string, error) {
	dbFile := profileDir + "/places.sqlite"
	info, err := os.Stat(dbFile)
	if os.IsNotExist(err) || info.IsDir() {
		return "", fmt.Errorf("could not find %q or it's a directory", dbFile)
	}
	return dbFile, nil
}