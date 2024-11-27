package userlogic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/yendelevium/mp3Downloader/internal/downloader"
	"github.com/yendelevium/mp3Downloader/internal/scraper"
)

func UserInput() (string, error) {
	// User logic
	fmt.Println("Enter the name of the song/artist")

	input := bufio.NewReader(os.Stdin)
	search, err := input.ReadString('\n')
	if err != nil {
		return "", err
	}
	// Changing input to match query-url
	search = strings.TrimSpace(search)
	re := regexp.MustCompile(" ")
	re2 := regexp.MustCompile("--")
	search = re.ReplaceAllString(search, "-")
	search = re2.ReplaceAllString(search, "-")
	return search, nil

}

func UserDownloadChoice(downloadLocation string, scrapedData []scraper.Data) error {
	// Menu for user to select the songs to be downloaded
	input := bufio.NewReader(os.Stdin)
	for key, elem := range scrapedData {
		fmt.Printf("\n%d Track Name: %s\nArtist Name:%s\n", key, elem.TrackName, elem.ArtistName)
	}
	fmt.Println("Enter the idx of the song you want to download(if multiple, separate by spaces)")
	indexes, err := input.ReadString('\n')
	if err != nil {
		return err
	}
	indexes = strings.TrimSpace(indexes)
	idx := strings.Split(indexes, " ")

	// Downloading selected songs
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
	return nil
}
