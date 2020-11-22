package mozurl

// Type MozHost wraps all URLs that share the same host
type MozHost struct {
	host       string
	urls       []MozURL
	visitCount int
}

func (mh *MozHost) HostName() string {
	return mh.host
}

func NewMozHost(host string, visitCount int) *MozHost {
	return &MozHost{
		host:       host,
		visitCount: visitCount,
	}
}

func (h *MozHost) URLs() []MozURL {
	return h.urls
}

func (h *MozHost) AddRawURL(rawurl string) {
	url := NewMozURL(rawurl, 0)
	h.urls = append(h.urls, url)
}
