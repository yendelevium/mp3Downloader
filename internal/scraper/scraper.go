package scraper

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

const protocol = "https://"
const domainName = ".hydr0.org"

type Data struct {
	ArtistName string
	TrackName  string
	Mp3URL     string
}

func Scrape(search string) []Data {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
	)

	scrapedData := make([]Data, 0)
	// Scraping the page to get the artist, song and URL to download the song
	c.OnHTML(".playlist-left", func(h *colly.HTMLElement) {
		artistName := h.ChildText(".playlist-name .playlist-name-artist .no-ajax")
		trackName := h.ChildText(".playlist-name .playlist-name-title .no-ajax")
		mp3URL := h.ChildAttr(".playlist-play.no-ajax", "data-url")

		// Removing '/' so it doesn't cause problems when creating file
		re := regexp.MustCompile("/")
		artistName = re.ReplaceAllString(artistName, " ")
		trackName = re.ReplaceAllString(trackName, " ")

		scrapedData = append(scrapedData, Data{
			ArtistName: artistName,
			TrackName:  trackName,
			Mp3URL:     mp3URL,
		})
	})

	if err := c.Visit(fmt.Sprintf("%s%s%s", protocol, search, domainName)); err != nil {
		log.Fatal(err)
	}
	return scrapedData
}
