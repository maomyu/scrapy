package handle

import (
	"fmt"
	"io"
	"github.com/PuerkitoBio/goquery"
	"github.com/yuwe1/scrapy/db"
	"strings"
)

type PoemInfo struct {
}
func getUrls(url string,size int)[]string{
	urls :=make([]string, 0)
	urlpl :=strings.Replace(url,"A1.aspx","A%d.aspx",1)
	for i :=1;i<=size;i++{
		urls = append(urls,fmt.Sprintf(urlpl,i))
		fmt.Println(fmt.Sprintf(urlpl,i))
	}
	return urls
}
func (p *PoemInfo)Worker(body io.Reader,url string){
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Errorf("doc err", err)
	}
	doc.Find(".cont").Each(func (i int,s *goquery.Selection){
		author := ""
		dynsty := ""
		content := ""
		title := ""
		title = strings.TrimSpace(s.Find("p").Find("a").Find("b").Text())
		authorAndDynsty := strings.TrimSpace(s.Find(".source").Text())
		authorAndDynstySlice :=  strings.Split(authorAndDynsty,"：")
		if len(authorAndDynstySlice)==2 {
			dynsty = authorAndDynstySlice[0]
			author = authorAndDynstySlice[1]
		}
		// fmt.Printf("作者：%s,朝代:%s,标题:%s\n",author,dynsty,title)
		s.Find(".contson").Each(func(i int, s *goquery.Selection) {
			content = strings.TrimSpace(s.Text())
			// fmt.Printf("内容：%s\n",content)
		})
	
		if author!="" && dynsty!="" && content!="" && title!="" {
			p := db.Poem{}
			p.Auther = author
			p.Title = title
			p.Content = content
			p.Dynasty = dynsty
			p.Save()
		}
	})
}