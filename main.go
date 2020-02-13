package main

import (
	"fmt"

	"github.com/yuwe1/scrapy/gofish"
	"github.com/yuwe1/scrapy/handle"
)

func main() {
	baseurl := "https://so.gushiwen.org/authors/"
	h := handle.AuthorHandle{}
	fish := gofish.NewGoFish()
	request, err := gofish.NewRequest("GET", baseurl, gofish.UserAgent, &h, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fish.Request = request
	fish.Visit()
}
