A small CLI to save bookmarks to SQLite. I want to use it to experiment with different SQL libraries. All SQL libraries should be used to implement `TaggedmarkService` , which the CLI app instantiates and uses.

# WTFDial Blog Notes

- https://www.gobeyond.dev/packages-as-layers/
- https://github.com/benbjohnson/wtf

Application types (really interfaces) in their own package - root package that other types depend on

```go
type DialService interface {
	FindDialByID(ctx context.Context, id int) (*Dial, error)
	FindDials(ctx context.Context, filter DialFilter) ([]*Dial, int, error)
	CreateDial(ctx context.Context, dial *Dial) error
	UpdateDial(ctx context.Context, id int, upd DialUpdate) (*Dial, error)
	DeleteDial(ctx context.Context, id int) error
}
```

The `wtf` package implements these servicces as interfaces. The `sqlite` package implemtents these interface. In `wtf`, the main struct instantiates `sqlite`'s implementation of the interfaces, then attaches them to `http`'s struct, then calls start on the http server. The `http` package *also* implements `wtf`'s interface as a client (by making requests to the server). 

The cmd/wtfd package instantiates an HTTP server (which does *NOT* fulfil the interfaces, but contains them) and shoves the `SQLite` implementation into it.

# Questions

- should my Tag struct have a reference to Taggedmark?

# CLI

```bash
taggedmarks 
	taggedmark
		create --url url --tag a --tag b
		read --url url # or --id id
		delete --url url # or --id id
		list --url --tag a
    import --type firefoxJSON --path bookmarks.json
	tag  # Let's ignore this for now :)
		rename --from old --from old2 --to new
		delete --name name
```



