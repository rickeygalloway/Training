// RTB: Real Time Bidding
package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	// We have 50 msec to return an answer
	// http://shipt.com, https://shipt.com
	url := "http://shipt.com"
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	fmt.Println(bidOn(ctx, url))
}

// If algo didn't finish in time, return a default bid
func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // Use buffered channel to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	//change sleep value d on line 46 to get different return values
	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid //if bestBid times out, this will be returned value
	}

}

var defaultBid = Bid{
	AdURL: "http://adsRus.com/default",
	Price: 0.03,
}

// Writting by Algo team, time to completion varies
func bestBid(url string) Bid {
	// Simulate work
	d := 500 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsRus.com/ad17",
		Price: 0.07,
	}
}

type Bid struct {
	AdURL string
	Price float64
}
