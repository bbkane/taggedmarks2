package main

import (
	"fmt"
	"os"

	"github.com/bbkane/warg"
	"github.com/bbkane/warg/command"
	"github.com/bbkane/warg/flag"
	"github.com/bbkane/warg/section"
	"github.com/bbkane/warg/value"

	"github.com/bbkane/taggedmarks2/moderncsqlitehandrolled"
)

func createTaggedmark(pf flag.PassedFlags) error {
	dbPath := pf["--db-path"].(string)

	ts, err := moderncsqlitehandrolled.NewTaggedmarkService(dbPath)
	if err != nil {
		return fmt.Errorf("db load errror: %w", err)
	}
	fmt.Println(ts)
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
						"Tag to add",
						value.String,
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
