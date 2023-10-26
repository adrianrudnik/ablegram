package indexer

type SearchOptions struct {
	PrimaryLanguage string
}

func NewSearchOptions() *SearchOptions {
	return &SearchOptions{
		PrimaryLanguage: "en",
	}
}
