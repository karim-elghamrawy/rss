package rss

import (
	"fmt"
	"testing"
)

var item = rss2_0Item{
	Link:       "http://www.marketwatch.com/story/dove-wants-girls-to-speakbeautiful-but-sparks-a-twitter-backlash-instead-2016-10-19?siteid=rss&utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A%20marketwatch%2Fpf%20(MarketWatch.com%20-%20Personal%20Finance%20News)",
	Enclosures: nil,
}

func TestGetImage(t *testing.T) {
	img := getImageRss2(item)
	fmt.Println(img)
}
