package main

import (
	"github.com/veganafro/trendly-backend/nyt"
)

func main() {
	nytClient := nyt.NewClient("testing")
	nytClient.GetMostShared(1)
}

