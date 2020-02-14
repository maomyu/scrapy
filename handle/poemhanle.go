package handle

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuwe1/scrapy/gofish"
	// "github.com/yuwe1/scrapy/gofish"
	// "github.com/yuwe1/scrapy/gofish"
)

type AuthorHandle struct {
}

var (
	baseurl = "https://so.gushiwen.org"
)

func (a *AuthorHandle) Worker(body io.Reader, url string) {

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Errorf("doc err", err)
	}
	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		// 输出内容
		// auther := s.Text()
		// fmt.Printf("%d auther=%s\n", i, auther)
		link, _ := s.Attr("href")
		h := PomeHomeHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseurl+link, gofish.UserAgent, &h, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()
	})
}

type PomeHomeHandle struct {
}

func (p *PomeHomeHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Errorf("doc err", err)
	}
	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
		// 获取连接
		link, _ := s.Attr("href")
		// fmt.Println("作品主页=", baseurl+link)
		pinfo := PoemInfo{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseurl+link, gofish.UserAgent, &pinfo, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()

	})
}
