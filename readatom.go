package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"
)

var formatsAtom = []string{time.RFC3339, time.RFC3339Nano}

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

func readatom(byteValue []byte) (title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText string) {

	replace := string(byteValue)

	replace = strings.Replace(replace, "&lt;", "<", -1)

	replace = strings.Replace(replace, "&gt;", ">", -1)

	byteValue = []byte(replace)

	var feed Feed
	xml.Unmarshal(byteValue, &feed)

	altText = feed.Entry[0].Summary[0].Img[0].Alt

	altText = strings.Replace(altText, "&quot;", "\"", -1)

	comicUrlImage = feed.Entry[0].Summary[0].Img[0].Src

	comicUrlImage = strings.Replace(comicUrlImage, ".png", "", 1)

	comicUrlImage = comicUrlImage + "_2x.png"

	comicUrl = feed.Entry[0].Id

	comicUrl = comicUrl[16:]
	comicUrl = "https://m.xkcd.com" + comicUrl
	fmt.Println(comicUrl)

	title = feed.Entry[0].Title

	d := feed.Entry[0].Updated

	lastBuildDate, _ := parseTimeAtom(d)

	lastBuildDateFormatted = lastBuildDate.Format("2006-01-02 03:04:05")

	return title, comicUrl, lastBuildDateFormatted, comicUrlImage, altText

}

func parseTimeAtom(input string) (time.Time, error) {
	for _, format := range formatsAtom {
		t, err := time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("Unrecognized time format")
}
