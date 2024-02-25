package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Id      string   `xml:"id"`
	Updated string   `xml:"updated"`
	Entry   []Entry  `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name  `xml:"entry"`
	Title   string    `xml:"title"`
	Link    string    `xml:"link"`
	Updated string    `xml:"updated"`
	Id      string    `xml:"id"`
	Summary []Summary `xml:"summary"`
}

type Summary struct {
	XMLName xml.Name `xml:"summary"`
	Img     []Img    `xml:"img"`
}

type Img struct {
	XMLName xml.Name `xml:"img"`
	Src     string   `xml:"src,attr"`
	Title   string   `xml:"title,attr"`
	Alt     string   `xml:"alt,attr"`
}

func readatom() (title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText string) {

	xmlFile, err := os.Open("atom.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	replace := string(byteValue)

	replace = strings.Replace(replace, "&lt;", "<", -1)

	replace = strings.Replace(replace, "&gt;", ">", -1)

	byteValue = []byte(replace)

	var feed Feed
	xml.Unmarshal(byteValue, &feed)

	altText = feed.Entry[0].Summary[0].Img[0].Alt

	altText = strings.Replace(altText, "&quot;", "\"", -1)

	comicUrlImage = feed.Entry[0].Summary[0].Img[0].Src

	comicUrl = feed.Entry[0].Id

	comicUrl = comicUrl[16:]
	comicUrl = "https://m.xkcd.com" + comicUrl
	fmt.Println(comicUrl)

	title = feed.Entry[0].Title

	d := feed.Entry[0].Updated

	lastBuildDate, err := time.Parse("2006-01-02T03:04:05Z", d)
	if err != nil {
		log.Fatal(err)
	}

	lastBuildDateFormatted = lastBuildDate.Format("2006-01-02 03:04:05")

	today := time.Now()

	todayFormatted := today.Format("2006-01-02 03:04:05")

	_ = todayFormatted

	// Re-enable this once testing is done
	err = os.Remove("atom.xml")
	if err != nil {
		fmt.Println(err)
	}

	return title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText

}
