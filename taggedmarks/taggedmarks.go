package taggedmarks

import (
	"context"
	"time"
)

type Tag struct {
	CreateTime time.Time
	ID         int
	Name       string
	UpdateTime time.Time
}

type Taggedmark struct {
	CreateTime time.Time
	ID         int
	UpdateTime time.Time
	URL        string
	Tags       []*Tag
}

type TaggedmarkUpdate struct {
	URL  string
	Tags []*Tag
}

type TaggedmarkQuery struct {
	URL string
}

type TaggedmarkService interface {
	// CreateTaggedmark places the passed taggedmark into the database.
	// Updates ID, CreateTime, and UpdateTime from the database
	CreateTaggedmark(ctx context.Context, tm *Taggedmark) error

	// ReadTaggedmark(ctx context.Context, id int) (*Taggedmark, error)
	// UpdateTaggedmark(ctx context.Context, id int, upd TaggedmarkUpdate) (*Taggedmark, error)
	// DeleteTaggedmark(ctx context.Context, id int) error
	// ListTaggedmarks(ctx context.Context, query TaggedmarkQuery) ([]*Taggedmark, error)
}
