package mozurl

import (
	"errors"
	"net/url"
	"strings"
)

// MozUrl is a representation of a row from `moz_places` but containing only the data we need
type MozUrl struct {
	rawurl     string
	visitCount int
}

func NewMozUrl(rawurl string, visitCount int) MozUrl {
	return MozUrl{
		rawurl:     rawurl,
		visitCount: visitCount,
	}
}

func (mu MozUrl) extractHost() string {
	u, _ := url.Parse(mu.rawurl)
	switch u.Scheme {
	case "about", "javascript":
		return u.Scheme + ":" + u.Opaque
	case "file":
		return u.Scheme + "://" + u.Host
	case "place":
		return parsePlacesUrl(u.Opaque)
	}
	return u.Host
}

func parsePlacesUrl(place string) string {
	return "place" + ":" + strings.Split(place, "=")[0]
}

// Domain wraps all URLs that share the same host
type Domain struct {
	host       string
	urls       []MozUrl
	visitCount int
}

func NewDomain(templateUrl MozUrl) *Domain {
	host := templateUrl.extractHost()
	return &Domain{
		host:       host,
		urls:       []MozUrl{templateUrl},
		visitCount: templateUrl.visitCount,
	}
}

func (d *Domain) Host() string {
	return d.host
}

func (d *Domain) CountURLs() int {
	return len(d.urls)
}

func (d *Domain) addUrl(url MozUrl) error {
	if url.extractHost() != d.host {
		return errors.New("New URL host does not match domain host")
	}
	d.urls = append(d.urls, url)
	d.visitCount += url.visitCount
	return nil
}

func (d *Domain) VisitCount() int {
	return d.visitCount
}

func (d *Domain) URLs() []MozUrl {
	return d.urls
}

// DomainStore packs all the available domain (or host).
type DomainStore struct {
	Domains map[string]*Domain
	hosts   []*Domain // hosts helps the domain store internally track all the domains it's holding to ease comparision during sorts
}

func NewDomainStore() *DomainStore {
	return &DomainStore{
		Domains: map[string]*Domain{},
		hosts:   []*Domain{},
	}
}

func (ds *DomainStore) PackFromURLs(urls []MozUrl) {
	for _, url := range urls {
		ds.AppendURL(url)
	}
}

func (ds *DomainStore) AppendURL(url MozUrl) {
	host := url.extractHost()
	_, exists := ds.Domains[host]
	if !exists {
		domain := NewDomain(url)
		ds.Domains[domain.host] = domain
		ds.hosts = append(ds.hosts, domain)
	} else {
		ds.Domains[host].addUrl(url)
	}
}

func (ds *DomainStore) SortedDomains() []*Domain {
	return ds.hosts
}

func (ds DomainStore) Len() int {
	return len(ds.hosts)
}

func (ds DomainStore) Less(i, j int) bool {
	return ds.hosts[j].VisitCount() < ds.hosts[i].VisitCount()
}

func (ds DomainStore) Swap(i, j int) {
	ds.hosts[i], ds.hosts[j] = ds.hosts[j], ds.hosts[i]
}
