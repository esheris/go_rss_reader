package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://www.wfh.io/categories/6/jobs.atom")
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if r.StatusCode != http.StatusOK {
		log.Fatal(r.Status)
	}

	var f Feed
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failure reading Body")
	}
	xerr := xml.Unmarshal(bodyBytes, &f)
	if xerr != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse XML feed (%v)\n", xerr)
	}
	fmt.Print(f)
	for _, t := range f.Entry {
		fmt.Printf(t.Title)
	}
}

type Feed struct {
	XMLName xml.Name
	ID      string  `xml:"id"`
	Link    string  `xml:"link"`
	Title   string  `xml:"title"`
	Updated string  `xml:"updated"`
	Entry   []Entry `xml:"entry"`
}

type Entry struct {
	XMLName   xml.Name `xml:"entry"`
	ID        string   `xml:"id"`
	Published string   `xml:"published"`
	Updated   string   `xml:"updated"`
	Link      string   `xml:"link"`
	Title     string   `xml:"title"`
	Content   string   `xml:"content"`
	Author    Person   `xml:"author"`
}

type Person struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
}
