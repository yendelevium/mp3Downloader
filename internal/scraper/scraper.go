package scraper

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
	"github.com/yendelevium/mp3Downloader/internal/downloader"
)

const protocol = "https://"
const domainName = ".hydr0.org"

func Scrape() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
	)
	wg := &sync.WaitGroup{}
	// Scraping the page to get the artist, song and URL to download the song
	c.OnHTML(".playlist-left",func(h *colly.HTMLElement) {
		artistName := h.ChildText(".playlist-name .playlist-name-artist .no-ajax")
		trackName := h.ChildText(".playlist-name .playlist-name-title .no-ajax")
		mp3URL := h.ChildAttr(".playlist-play.no-ajax","data-url")
		fmt.Println(trackName,"-",artistName)

		// Concurrent downloads so its faster
		wg.Add(1)
		go downloader.DownloadSong(trackName,artistName,mp3URL,wg)
	})

	if err:=c.Visit(fmt.Sprintf("%sharehare-ya%s",protocol,domainName));err!=nil{
		log.Fatal(err)
	}
	wg.Wait()
}