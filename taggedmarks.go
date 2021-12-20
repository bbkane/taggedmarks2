package taggedmarks2

import "context"

type Tag struct {
	ID   int
	Name string
}

type Taggedmark struct {
	ID   int
	URL  string
	Tags []Tag
}

type TaggedmarkUpdate struct {
	URL  string
	Tags []Tag
}

type TaggedmarkQuery struct {
	URL string
}

type TaggedmarkService interface {
	CreateTaggedmark(ctx context.Context, tm *Taggedmark) error
	// ReadTaggedmark(ctx context.Context, id int) (*Taggedmark, error)
	// UpdateTaggedmark(ctx context.Context, id int, upd TaggedmarkUpdate) (*Taggedmark, error)
	// DeleteTaggedmark(ctx context.Context, id int) error
	// ListTaggedmarks(ctx context.Context, query TaggedmarkQuery) ([]*Taggedmark, error)
}
