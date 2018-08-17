package serarch

import (
	"log"
	//"fmt"
)

type Result struct {
	Filed   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Macth(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {

	serarchResults, err := matcher.Search(feed, searchTerm)

	if err != nil {

		log.Println(err)
		return
	}
	for _, result := range serarchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result :=results{
		fmt.Println("%s:\n%s\n\n",result.Filed,result.Content)
	}
}
