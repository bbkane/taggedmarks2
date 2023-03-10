package main

import (
	"context"
	"fmt"
	"time"

	"go.bbkane.com/warg/command"

	"go.bbkane.com/taggedmarks2/moderncsqlite"
	"go.bbkane.com/taggedmarks2/taggedmarks"
)

func createTaggedmark(pf command.Context) error {
	dbPath := pf.Flags["--db-path"].(string)
	url := pf.Flags["--url"].(string)
	sqlBackend := pf.Flags["--sql-backend"].(string)
	tagsFlag := []string{}
	if tagsF, exists := pf.Flags["--tag"]; exists {
		tagsFlag = tagsF.([]string)
	}

	var ts taggedmarks.TaggedmarkService
	var tsErr error
	switch sqlBackend {
	case "moderncsqlite":
		ts, tsErr = moderncsqlite.NewTaggedmarkService(dbPath)
	default:
		return fmt.Errorf("Unknown sql backend: %s", sqlBackend)
	}

	if tsErr != nil {
		return fmt.Errorf("db load errror: %w", tsErr)
	}

	now := time.Now()
	tags := []*taggedmarks.Tag{}
	for _, t := range tagsFlag {
		tags = append(
			tags,
			&taggedmarks.Tag{
				Name:       t,
				CreateTime: now,
				UpdateTime: now,
				ID:         0,
			},
		)
	}

	tm := &taggedmarks.Taggedmark{
		URL:        url,
		Tags:       tags,
		CreateTime: now,
		UpdateTime: now,
		ID:         0,
	}

	err := ts.CreateTaggedmark(context.Background(), tm)
	if err != nil {
		err = fmt.Errorf("createTaggedmark err: %w", err)
		return err
	}

	fmt.Printf("%#v\n", *tm)
	return nil
}
