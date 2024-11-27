package main

import (
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/yendelevium/mp3Downloader/internal/scraper"
	userlogic "github.com/yendelevium/mp3Downloader/internal/userLogic"
)

func main() {
	// Loading the env file to get the download location for the songs
	// since .env is in the root directory, write the appropriate path
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Counldn't find .env file, defaulting to downloading in cmd/mp3Downloader %s", err)

	}

	// Getting directory path(based on OS used, as WSL and Windows have different paths)
	// If .env file was not found, the download location by default will be the cwd
	var downloadLocation string
	if runtime.GOOS == "windows" {
		downloadLocation = os.Getenv("downloadLocationWindows")
	} else {
		downloadLocation = os.Getenv("downloadLocationUnix")
	}

	search, err := userlogic.UserInput()
	if err != nil {
		log.Fatal(err)
	}
	// Scraping the web for the given song/artist
	scrapedData, err := scraper.Scrape(search)
	if err != nil {
		log.Fatal(err)
	}

	err = userlogic.UserDownloadChoice(downloadLocation, scrapedData)
	if err != nil {
		log.Fatal(err)
	}

}
