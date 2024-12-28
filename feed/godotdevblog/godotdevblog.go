package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	http "net/http"
	"time"
)

type AtomFeed struct {
	Title   string      `xml:"title"`
	Id      string      `xml:"id"`
	Link    Link        `xml:"link"`
	Updated string      `xml:"updated"`
	Entry   []AtomEntry `xml:"entry"`
}

type AtomEntry struct {
	Title     string   `xml:"title"`
	Id        string   `xml:"id"`
	Link      Link     `xml:"link"`
	Published string   `xml:"published"`
	Updated   string   `xml:"updated"`
	Author    []Author `xml:"author"`
	Summary   Summary  `xml:"summary"`
	Content   Content  `xml:"content"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Author struct {
	Name string `xml:"name"`
}

type Summary struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Content struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

func main() {

	url := "https://go.dev/blog/feed.atom"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := &AtomFeed{}
	err = xml.Unmarshal(body, feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range feed.Entry {
		fmt.Println("Title:", entry.Title)
		fmt.Print("Authors: ")
		for _, author := range entry.Author {
			fmt.Print(author.Name + " ")
		}
		fmt.Println()

		t, _ := time.Parse(time.RFC3339, entry.Published)

		fmt.Println("Published:", t.Format(time.DateOnly))
		fmt.Println("Link:", entry.Link.Href)
		fmt.Println()
	}
}
