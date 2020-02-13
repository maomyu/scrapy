package handle

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type AuthorHandle struct {
}

var (
	baseurl = "https://so.gushiwen.org/authors/"
)

func (a *AuthorHandle) Worker(body io.Reader, url string) {

	doc, err := goquery.NewDocumentFromReader(body)
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
