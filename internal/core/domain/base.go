package domain

import (
	"fmt"
	"net/http"
)

type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type Page struct {
	Count      int
	PageSize   int
	PageNumber int
	Links      *Links `json:"links"`
}

func (page *Page) SetLinks(r *http.Request) {
	page.Links = &Links{
		Self: page.GetSelfLink(r),
		Next: page.GetNextLink(r),
		Prev: page.GetPrevLink(r),
	}
}

func (page *Page) GetNextLink(r *http.Request) string {
	if page.PageNumber >= page.GetMaxPage() {
		return ""
	}

	return formatLink(r, page.PageNumber+1, page.PageSize)
}

func (page *Page) GetSelfLink(r *http.Request) string {
	return formatLink(r, page.PageNumber, page.PageSize)
}

func (page *Page) GetPrevLink(r *http.Request) string {
	if page.PageNumber < 2 || (page.PageNumber-1) > page.GetMaxPage() {
		return ""
	}

	return formatLink(r, page.PageNumber-1, page.PageSize)
}

func (page *Page) GetMaxPage() int {
	return page.Count / page.PageSize
}

func formatLink(r *http.Request, pageNumber int, pageSize int) string {
	return fmt.Sprintf("%s%s?pageNumber=%d&pageSize=%d", r.Host, r.URL.Path, pageNumber, pageSize)
}
