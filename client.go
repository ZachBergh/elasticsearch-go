package elasticsearchgo

type EsClient struct {
	Domain string
	Index  string
	Type   string
	Query  []byte
}

func NewClient(domain, index, itype, query string) *EsClient {
	return &EsClient{domain, index, itype, []byte(query)}
}
