package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	authers := "https://so.gushiwen.org/authors/"
	res, err := http.Get(authers)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Errorf("doc err", err)
	}
	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		// 输出内容
		auther := s.Text()
		fmt.Printf("%d auther=%s\n", i, auther)
		link, _ := s.Attr("href")
		fmt.Println(i, "link = ", link)
	})

}
