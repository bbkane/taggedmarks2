package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bbkane/warg"
	"github.com/bbkane/warg/command"
	"github.com/bbkane/warg/flag"
	"github.com/bbkane/warg/section"
	"github.com/bbkane/warg/value"

	"github.com/bbkane/taggedmarks2"
	"github.com/bbkane/taggedmarks2/moderncsqlitehandrolled"
)

func createTaggedmark(pf flag.PassedFlags) error {
	dbPath := pf["--db-path"].(string)
	url := pf["--url"].(string)
	tagsFlag := []string{}
	if tagsF, exists := pf["--tag"]; exists {
		tagsFlag = tagsF.([]string)
	}

	ts, err := moderncsqlitehandrolled.NewTaggedmarkService(dbPath)
	if err != nil {
		return fmt.Errorf("db load errror: %w", err)
	}

	tags := []*taggedmarks2.Tag{}
	for _, t := range tagsFlag {
		tags = append(tags, &taggedmarks2.Tag{Name: t})
	}

	tm := &taggedmarks2.Taggedmark{
		URL:  url,
		Tags: tags,
	}

	err = ts.CreateTaggedmark(context.Background(), tm)
	if err != nil {
		err = fmt.Errorf("createTaggedmark err: %w", err)
		return err
	}

	fmt.Println(*tm)
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
				value.Path,
				flag.Default("taggedmarks2.db"),
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
						value.StringSlice,
					),
					command.Flag(
						"--url",
						"URL",
						value.String,
						flag.Required(),
					),
				),
			),
		),
	)

	app.MustRun(os.Args, os.LookupEnv)
}
