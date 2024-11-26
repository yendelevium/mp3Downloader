package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// Filenm will be Tracknm-Artistnm.mp3
func DownloadSong(trackName string, artistName string,mp3URL string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(mp3URL)
	if err != nil {
		log.Fatal("Get Req", err)
	}

	// TODO: Designated directory to download songs?
	// TODO: Need to strip track and artist name from '/' as os.Create will think of it as a directory
	file, err := os.Create(fmt.Sprintf("%s-%s%s", trackName, artistName, ".mp3"))
	if err != nil {
		log.Fatal("File Creation:", err)
	}

	// Downloading the mp3 by copying the file from resp.Body
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Writing:", err)
	}
	log.Println("Downloaded a song")

}