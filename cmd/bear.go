package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
)

func getBearDBPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dbpath := fmt.Sprintf("%s/Library/Containers/net.shinyfrog.bear/Data/Documents/Application Data/database.sqlite", usr.HomeDir)

	info, err := os.Stat(dbpath)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		return "", fmt.Errorf("Bear db [%s] is not a regular file", dbpath)
	}

	return dbpath, nil
}

func getNotesFromDB(db *sql.DB) ([]*Note, error) {

	rows, err := db.Query("SELECT ZTITLE, ZTEXT, ZTRASHED, ZCREATIONDATE FROM ZSFNOTE;")
	if err != nil {
		return nil, err
	}

	notes := []*Note{}

	for rows.Next() {
		entry := &Note{}

		err = rows.Scan(&entry.Title, &entry.FullText, &entry.Trashed, &entry.CreatedOn)
		if err != nil {
			return nil, err
		}

		if entry.Deleted() {
			continue
		}

		err = entry.Load()
		if err != nil {
			return nil, err
		}

		notes = append(notes, entry)
	}

	return notes, nil
}
