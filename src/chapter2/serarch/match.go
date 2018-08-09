package serarch

import (
	"log"
)

type Result struct {
	Filed   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
