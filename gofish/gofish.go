package gofish

import "time"

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"
	// 限流
	Qps = 50
)

var rateLimter = time.Tick(time.Second / Qps)

type GoFish struct {
	Request *Request
}

func NewGoFish() *GoFish {
	return &GoFish{}
}

func (g *GoFish) Visit() error {
	return g.Request.Do()
}
