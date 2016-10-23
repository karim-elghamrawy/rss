package rss

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
)

func getImageRss1(item rss1_0Item) string {
	for _, enc := range item.Enclosures {
		if strings.Contains(enc.Type, "image") {
			return enc.Url
		}
	}

	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}

	selection := doc.Find(`meta[property="og:image"]`)
	if selection != nil {
		if imgURL, found := selection.Attr("content"); found {
			return imgURL
		}
	}

	selection = doc.Find(`meta[property="twitter:image"]`)
	if selection != nil {
		if imgURL, found := selection.Attr("content"); found {
			return imgURL
		}
	}

	return ""
}

func getSummaryRss1(item rss1_0Item) string {
	var fbDesc, twDesc, socDesc, rssDesc string
	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}
	selection := doc.Find(`meta[property="og:description"]`)
	if selection != nil {
		if desc, found := selection.Attr("content"); found {
			fbDesc = desc
		}
	}

	selection = doc.Find(`meta[property="twitter:description"]`)
	if selection != nil {
		if desc, found := selection.Attr("content"); found {
			twDesc = desc
		}
	}

	if !strings.ContainsAny(item.Description, "&lt;&gt;") {
		rssDesc = item.Description
	}

	if len(fbDesc) >= len(twDesc) {
		socDesc = fbDesc
	} else {
		socDesc = twDesc
	}

	if len(socDesc) >= len(rssDesc) {
		return socDesc
	}
	return rssDesc
}

func getImageRss2(item rss2_0Item) string {
	for _, enc := range item.Enclosures {
		if strings.Contains(enc.Type, "image") {
			return enc.Url
		}
	}
	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}

	selection := doc.Find(`meta[property="og:image"]`)
	if selection != nil {
		if imgURL, found := selection.Attr("content"); found {
			return imgURL
		}
	}

	selection = doc.Find(`meta[property="twitter:image"]`)
	if selection != nil {
		if imgURL, found := selection.Attr("content"); found {
			return imgURL
		}
	}

	return ""
}

func getSummaryRss2(item rss2_0Item) string {
	var fbDesc, twDesc, socDesc, rssDesc string
	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}
	selection := doc.Find(`meta[property="og:description"]`)
	if selection != nil {
		if desc, found := selection.Attr("content"); found {
			fbDesc = desc
		}
	}

	selection = doc.Find(`meta[property="twitter:description"]`)
	if selection != nil {
		if desc, found := selection.Attr("content"); found {
			twDesc = desc
		}
	}

	if !strings.ContainsAny(item.Description, "&lt;&gt;") {
		rssDesc = item.Description
	}

	if len(fbDesc) >= len(twDesc) {
		socDesc = fbDesc
	} else {
		socDesc = twDesc
	}

	if len(socDesc) >= len(rssDesc) {
		return socDesc
	}
	return rssDesc
}
