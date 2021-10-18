package domain

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
