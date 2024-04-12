package db

func Populate[T IsEmptier](tags []string, store Storer[T]) ([]T, error) {
	docs := make([]T, 0)
	for _, v := range tags {
		doc, err := store.One(Query{"tag": v})
		if err != nil {
			return nil, err
		}
		docs = append(docs, *doc)
	}
	return docs, nil
}
