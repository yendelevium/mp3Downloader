package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/yendelevium/mp3Downloader/internal/downloader"
	"github.com/yendelevium/mp3Downloader/internal/scraper"
)

func main() {
	// Loading the env file to get the download location for the songs
	// It has to be an absolute path
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Getting directory path(based on OS used, as WSL and Windows have different paths)
	var downloadLocation string
	if runtime.GOOS == "windows" {
		downloadLocation = os.Getenv("downloadLocationWindows")
	} else {
		downloadLocation = os.Getenv("downloadLocationUnix")
	}
	// User Logic
	fmt.Println("Enter the name of the song/artist")
	input := bufio.NewReader(os.Stdin)
	search, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	search = strings.TrimSpace(search)
	re := regexp.MustCompile(" ")
	re2 := regexp.MustCompile("--")
	search = re.ReplaceAllString(search, "-")
	search = re2.ReplaceAllString(search, "-")

	// Scraping the web for the given song/artist
	scrapedData := scraper.Scrape(search)

	// User Logic
	for key, elem := range scrapedData {
		fmt.Printf("\n%d Track Name: %s\nArtist Name:%s\n", key, elem.TrackName, elem.ArtistName)
	}
	fmt.Println("Enter the idx of the song you want to download(if multiple, separate by spaces)")
	indexes, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	indexes = strings.TrimSpace(indexes)
	idx := strings.Split(indexes, " ")

	wg := &sync.WaitGroup{}

	fmt.Println("Starting Download")
	for _, elem := range idx {
		songIdx, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		if songIdx < int64(len(scrapedData)) && songIdx > 0 {
			trackName := scrapedData[songIdx].TrackName
			artistName := scrapedData[songIdx].ArtistName
			mp3URL := scrapedData[songIdx].Mp3URL

			// Concurrent downloads for faster downloading
			wg.Add(1)
			go downloader.DownloadSong(trackName, artistName, mp3URL, downloadLocation, wg)

		} else {
			log.Printf("Idx out of range%d", songIdx)
		}
	}

	wg.Wait()
	fmt.Println("Finished Downloading")

}
