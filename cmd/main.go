package main

import (
	"fmt"

	"gitlab.com/idoko/foxtop/mozurl"
)

func main() {
	raw := map[string]int{
		"https://t.co/lorem":              4,
		"https://twitter.com/rxrog/":      2,
		"https://twitter.com/theshalvah/": 1,
		"https://mchl.xyz":                3,
	}
	mozUrls := []mozurl.MozUrl{}
	for u, c := range raw {
		url := mozurl.NewMozUrl(u, c)
		mozUrls = append(mozUrls, url)
	}

	store := mozurl.NewDomainStore()
	store.PackFromURLs(mozUrls)

	/*store.Sort()*/
	for _, domain := range store.Domains {
		fmt.Println(domain.Host())
	}
}
