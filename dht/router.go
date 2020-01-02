package dht

type Route struct {
	buckets []*Bucket
}

func NewRoute() *Route {
	return &Route{buckets:make([]*Bucket,8)}
}