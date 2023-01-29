package main

import (
	"os"

	"go.bbkane.com/warg"
	"go.bbkane.com/warg/command"
	"go.bbkane.com/warg/flag"
	"go.bbkane.com/warg/section"
	"go.bbkane.com/warg/value/scalar"
	"go.bbkane.com/warg/value/slice"
)

var version string

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
