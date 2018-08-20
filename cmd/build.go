package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	// sqlite3 imported for bear db
	_ "github.com/mattn/go-sqlite3"
)

var workdir string

// BuildOptions ...
type BuildOptions struct {
	ContentDirectory        string
	ContentSectionDirectory string
	BearSQLiteDBPath        string
	UseHugoExec             bool
}

// Build ...
func Build(opts *BuildOptions) error {
	var err error

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	workdir = cwd

	if opts.BearSQLiteDBPath == "" {
		opts.BearSQLiteDBPath, err = getBearDBPath()
		if err != nil {
			return err
		}
	}

	db, err := sql.Open("sqlite3", opts.BearSQLiteDBPath)
	if err != nil {
		return err
	}

	notes, err := getNotesFromDB(db)
	if err != nil {
		return err
	}

	color.Cyan("Found %d notes\n", len(notes))

	contentdir := filepath.Join(cwd, opts.ContentDirectory)
	err = os.MkdirAll(contentdir, 0770)
	if err != nil {
		return err
	}

	contentsection := filepath.Join(contentdir, opts.ContentSectionDirectory)
	err = os.RemoveAll(contentsection)
	if err != nil {
		return err
	}

	err = os.MkdirAll(contentsection, 0770)
	if err != nil {
		return err
	}

	for _, note := range notes {
		title := note.Title
		if title == "" {
			ln := len(note.FullText)
			if ln >= 20 {
				ln = 20
			}

			color.Yellow("warning: skipping note - does not have a title [%s]\n", strings.Replace(note.FullText[:ln], "\n", " ", -1))
			continue
		}

		title = strings.TrimSuffix(
			strings.Replace(
				strings.Replace(
					strings.Replace(
						strings.ToLower(title), " - ", "-", -1), "/", "-", -1), " ", "-", -1), ".md")

		path := filepath.Join(contentsection, title+".md")

		d, err := note.Marshal()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path, d, 0660)
		if err != nil {
			return err
		}
	}

	// files have been written, time for hugo
	hugo, err := NewHugoSite(workdir)
	if err != nil {
		return err
	}

	color.Cyan("Building Hugo site at %s\n", workdir)
	return hugo.Build()
}

func jsonmi(v interface{}) {
	d, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("jsonmi start -->\n%s\n<-- jsonmi end\n", string(d))
}
