package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Rss     Channel  `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Atom          string   `xml:"atom:link"`
	Language      string   `xml:"language"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Item          []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
}

func readxml() (title, comicUrl, lastBuildDateFormatted string) {

	xmlFile, err := os.Open("current.rss")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	var rss Rss
	xml.Unmarshal(byteValue, &rss)

	comicUrl = rss.Rss.Item[0].Link

	title = rss.Rss.Item[0].Title

	d := rss.Rss.LastBuildDate

	lastBuildDate, err := time.Parse("Mon, 2 Jan 2006 03:04:05 -0700", d)
	if err != nil {
		log.Fatal(err)
	}

	lastBuildDateFormatted = lastBuildDate.Format("2006-01-02 03:04:05")

	today := time.Now()

	todayFormatted := today.Format("2006-01-02 03:04:05")

	_ = todayFormatted

	// Re-enable this once testing is done
	err = os.Remove("current.rss")
	if err != nil {
		fmt.Println(err)
	}

	return title, comicUrl, lastBuildDateFormatted

}
