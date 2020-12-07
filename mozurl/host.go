package mozurl

// Type MozHost wraps all URLs that share the same Host
type MozHost struct {
	Host       string   `json:"Host"`
	URLs       []MozURL `json:"URLs"`
	VisitCount int      `json:"visit_count"`
}

func (mh *MozHost) HostName() string {
	return mh.Host
}

func NewMozHost(host string, visitCount int) *MozHost {
	return &MozHost{
		Host:       host,
		VisitCount: visitCount,
	}
}

func (h *MozHost) AddRawURL(rawurl string) {
	url := NewMozURL(rawurl, 0)
	h.URLs = append(h.URLs, url)
}
