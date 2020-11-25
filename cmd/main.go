package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"gitlab.com/idoko/foxtop"
)

func main() {
	var dbFile string
	var err error
	var profileDir string

	flag.StringVarP(&profileDir, "profile-path", "p", "", "firefox profile path")
	flag.Parse()

	if profileDir == "" {
		if profileDir, err = defaultProfileDir(); err != nil {
			log.Fatal(err)
		}
	}
	dbFile, err = getDbFile(profileDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dbFile)
}

func getDbFile(profileDir string) (string, error) {
	dbFile := profileDir + "/places.sqlite"
	info, err := os.Stat(dbFile)
	if os.IsNotExist(err) || info.IsDir() {
		return "", fmt.Errorf("could not find %q or it's a directory", dbFile)
	}
	return dbFile, nil
}

func defaultProfileDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	firefoxPath := home + "/.mozilla/firefox"

	pfConfig, err := os.Open(firefoxPath + "/profiles.ini")
	if err != nil {
		return "", err
	}
	cfg, err := foxtop.LoadConfig(pfConfig)
	if err != nil {
		return "", err
	}
	return cfg.DefaultPath(), nil
}
