Let's try to do taggedmarks again, but following https://www.gobeyond.dev/wtf-dial/ and https://github.com/benbjohnson/wtf .

I want to abstract the storage (and then use that to try different DB frameworks). I also want to use warg :D

# SQL Things to Try

- hand rolled
- Ent
- https://github.com/Masterminds/squirrel
  - I don't think this support SQLite3
- https://github.com/upper/db (need to fork to use modernc/sqlite?)

# Other things to Try

- GraphQL with Ent? It would be nice generate the graphql code

# Error Handling

- https://www.gobeyond.dev/failure-is-your-domain/
  - I don't really like this

# WTFDial Blog Notes

https://www.gobeyond.dev/packages-as-layers/

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

# TODO

Let's start with the interface idea. We want to CRUDL taggedmarks, and we can make interfaces for that, then make a hand-rolled SQL implementation of that and a cli package to drive it... but first, lunch!

- pass times into CreateDial instead of generating them there
- move cmd/taggedmarks2 into root, and root into new package model
- Get auth working, errors working, metrics?
- get HTTP working

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



