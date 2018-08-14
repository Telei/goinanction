package serarch

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher

	Register("default", matcher)

}

func (m,defaultMatcher) Search (feed *Feed,searchTerm string)([]*Result,error) {
	retrun nil,nil;
	
}
