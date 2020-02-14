package handle

import (
	"testing"

	"github.com/yuwe1/scrapy/db"
)

func Test_CreateImg(t *testing.T) {
	poems, err := db.QueryPoemsByAuthor("王安石")
	if err != nil {
		return
	}
	for index := range poems {
		CreateShiImage(poems[index])
	}
}
