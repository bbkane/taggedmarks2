package main

import (
	"context"
	"fmt"
	"os"

	"go.bbkane.com/warg"
	"go.bbkane.com/warg/command"
	"go.bbkane.com/warg/flag"
	"go.bbkane.com/warg/section"
	"go.bbkane.com/warg/value/scalar"
	"go.bbkane.com/warg/value/slice"

	"go.bbkane.com/taggedmarks2/moderncsqlite"
	"go.bbkane.com/taggedmarks2/taggedmarks"
)

var version string

func createTaggedmark(pf command.Context) error {
	dbPath := pf.Flags["--db-path"].(string)
	url := pf.Flags["--url"].(string)
	tagsFlag := []string{}
	if tagsF, exists := pf.Flags["--tag"]; exists {
		tagsFlag = tagsF.([]string)
	}

	var ts taggedmarks.TaggedmarkService
	ts, err := moderncsqlite.NewTaggedmarkService(dbPath)
	if err != nil {
		return fmt.Errorf("db load errror: %w", err)
	}

	tags := []*taggedmarks.Tag{}
	for _, t := range tagsFlag {
		tags = append(tags, &taggedmarks.Tag{Name: t})
	}

	tm := &taggedmarks.Taggedmark{
		URL:  url,
		Tags: tags,
	}

	err = ts.CreateTaggedmark(context.Background(), tm)
	if err != nil {
		err = fmt.Errorf("createTaggedmark err: %w", err)
		return err
	}

	fmt.Printf("%#v\n", *tm)
	return nil
}

func main() {
	app := warg.New(
		"taggedmarks2",
		section.New(
			"Save bookmarks with tags!",
			section.Flag(
				"--db-path",
				"Path to sqlite DB",
				scalar.Path(
					scalar.Default("taggedmarks2.db"),
				),
				flag.Required(),
			),
			section.Section(
				"taggedmark",
				"Taggedmark commands",
				section.Command(
					"create",
					"Create a taggedmark",
					createTaggedmark,
					command.Flag(
						"--tag",
						"Tags to add",
						slice.String(),
					),
					command.Flag(
						"--url",
						"URL",
						scalar.String(),
						flag.Required(),
					),
				),
			),
		),
		warg.AddColorFlag(),
		warg.AddVersionCommand(version),
	)

	app.MustRun(os.Args, os.LookupEnv)
}
