package serarch

import (
	"log"
	"sync"
)

//注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

//Run执行搜索的逻辑
func Run(searchTerm string) {

	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists = matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(macher, feed)

	}

	go func() {
		waitGroup.Wait()

		close(results)
	}()

	Display(results)
}
