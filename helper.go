package rss

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
)

const (
	// MaxDescriptionWords is the maximum number of words in a description
	MaxDescriptionWords = 25
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

	var desc string

	if len(socDesc) >= len(rssDesc) {
		desc = socDesc
	} else {
		desc = rssDesc
	}

	descSlice := strings.Split(desc, " ")

	if len(descSlice) > MaxDescriptionWords {
		desc = strings.Join(descSlice[:MaxDescriptionWords], " ")
	}
	return desc
}

func getTitleRss1(item rss1_0Item) string {
	var fbTitle, twTitle, socTitle, rssTitle string
	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}
	selection := doc.Find(`meta[property="og:title"]`)
	if selection != nil {
		if t, found := selection.Attr("content"); found {
			fbTitle = t
		}
	}

	selection = doc.Find(`meta[property="twitter:title"]`)
	if selection != nil {
		if t, found := selection.Attr("content"); found {
			twTitle = t
		}
	}

	rssTitle = item.Title

	if len(fbTitle) >= len(twTitle) {
		socTitle = fbTitle
	} else {
		socTitle = twTitle
	}

	var title string

	if len(socTitle) >= len(rssTitle) {
		title = socTitle
	} else {
		title = rssTitle
	}
	return title
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

	var desc string

	if len(socDesc) >= len(rssDesc) {
		desc = socDesc
	} else {
		desc = rssDesc
	}

	descSlice := strings.Split(desc, " ")

	if len(descSlice) > MaxDescriptionWords {
		desc = strings.Join(descSlice[:MaxDescriptionWords], " ")
	}
	return desc
}

func getTitleRss2(item rss2_0Item) string {
	var fbTitle, twTitle, socTitle, rssTitle string
	doc, err := goquery.NewDocument(item.Link)
	if err != nil {
		log.Warnf("can not parse url %v", item.Link)
		return ""
	}
	selection := doc.Find(`meta[property="og:title"]`)
	if selection != nil {
		if t, found := selection.Attr("content"); found {
			fbTitle = t
		}
	}

	selection = doc.Find(`meta[property="twitter:title"]`)
	if selection != nil {
		if t, found := selection.Attr("content"); found {
			twTitle = t
		}
	}

	rssTitle = item.Title

	if len(fbTitle) >= len(twTitle) {
		socTitle = fbTitle
	} else {
		socTitle = twTitle
	}

	var title string

	if len(socTitle) >= len(rssTitle) {
		title = socTitle
	} else {
		title = rssTitle
	}
	return title
}
