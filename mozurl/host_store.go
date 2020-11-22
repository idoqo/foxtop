package mozurl

// DomainStore packs all the available domain (or host).
/*type DomainStore struct {
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
}*/
