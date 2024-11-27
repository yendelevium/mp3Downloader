package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Filenm will be Tracknm-Artistnm.mp3
func DownloadSong(trackName string, artistName string, mp3URL string, downloadLocation string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(mp3URL)
	if err != nil {
		log.Fatal("Get Req", err)
	}

	err = os.MkdirAll(downloadLocation, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directories: %s", err)
	}
	path := filepath.Join(downloadLocation, fmt.Sprintf("%s-%s%s", trackName, artistName, ".mp3"))
	path = filepath.FromSlash(path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("File Creation:", err)
	}

	// Downloading the mp3 by copying the file from resp.Body
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Writing:", err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal("Closing:", err)
	}

	log.Printf("Downloaded %s", fmt.Sprintf("%s-%s%s", trackName, artistName, ".mp3"))

}
