package main

func main() {
	/*raw := map[string]int{
		"https://t.co/lorem":         45,
		"https://twitter.com/rxrog/": 16,
	}
	mozUrls := []mozurl.MozUrl{}
	for u, c := range raw {
		url := mozurl.NewMozUrl(u, c)
		mozUrls = append(mozUrls, url)
	}
	dbFile := "tmp.sqlite"
	db, err := db.Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	mozUrls, err := db.GetURLs()
	if err != nil {
		log.Fatal(err)
	}

	store := mozurl.NewDomainStore()
	store.PackFromURLs(mozUrls)

	sort.Sort(store)
	for _, domain := range store.SortedDomains() {
		fmt.Println(domain.Host())
	}*/
}
