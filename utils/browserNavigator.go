package utils

import (
	"regexp"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var (
	kagiUrl = "https://kagi.com/search?token=OAGAXyqQ6wc.si8ThaxViKOFAq7t7OvK8vRUswgoiVIbsiXrL5P1Pdk&q="
)

type BrowserNavigator struct {
	browser *rod.Browser
}

func (b *BrowserNavigator) cleanupHtml(html string) string {
	matcher := regexp.MustCompile(`<svg[^>]*>(?:<.*?>|\s)*?</svg>|<script(?:.|\n)*script>|(?:class|style)="[^\"]*"|<!--.*?-->`)
	stripOutSpaces := regexp.MustCompile(`\s+`)
	removeExtraSpaces := regexp.MustCompile(`>\s+<`)

	strippedOut := matcher.ReplaceAllLiteralString(html, "")
	strippedOut = stripOutSpaces.ReplaceAllLiteralString(strippedOut, " ")
	return removeExtraSpaces.ReplaceAllLiteralString(strippedOut, "><")

}

func (b *BrowserNavigator) GetContentsFromWebpage(url string, selector *string) *string {
	openedPage := b.browser.MustPage(url).MustWaitIdle().MustWaitDOMStable()
	defer b.browser.Close()

	elemToSelect := "body"
	if selector != nil {
		elemToSelect = *selector
	}

	selected := openedPage.MustElement(elemToSelect).MustHTML()

	out := b.cleanupHtml(selected)

	return &out
}

func (b *BrowserNavigator) PerformSearch(searchTerm string) *string {
	fullUrl := kagiUrl + searchTerm

	return b.GetContentsFromWebpage(fullUrl, StrPnt("#main"))
}

func NewBrowserNavigator() *BrowserNavigator {
	u := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(u).Trace(true).SlowMotion(time.Second).MustConnect()
	return &BrowserNavigator{
		browser,
	}
}
