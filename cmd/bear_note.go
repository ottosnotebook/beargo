package cmd

import (
	"bytes"
	"time"

	"github.com/go-yaml/yaml"
)

// Note represents a bear note.
type Note struct {
	Title     string
	CreatedOn float64
	FullText  string
	Trashed   int

	Frontmatter *Frontmatter
	Body        []byte
}

// Frontmatter ...
type Frontmatter map[string]interface{}

// Load ...
func (b *Note) Load() error {
	var err error

	b.Frontmatter = &Frontmatter{
		"title": b.Title,
		"date":  time.Unix(int64(b.CreatedOn), 0).Format("2006-01-02"),
	}

	bod := b.FullText

	bod, err = scanFor(bod, "[image:", "]")
	if err != nil {
		return err
	}

	bod, err = scanFor(bod, "[file:", "]")
	if err != nil {
		return err
	}

	b.Body = []byte(bod)
	return nil
}

func scanFor(body string, prefix, suffix string) (string, error) {

	var (
		r   = bytes.NewBuffer([]byte(body))
		buf = bytes.NewBuffer([]byte{})

		start int
		c     byte
		err   error
	)

	for {

		if r.Len() <= 0 {
			return buf.String(), nil
		}

		if c, err = r.ReadByte(); err != nil {
			return "", err
		}

		if err := buf.WriteByte(c); err != nil {
			return "", err
		}

		if bytes.HasSuffix(buf.Bytes(), []byte(prefix)) {
			start = buf.Len() - len(prefix)
			continue
		}

		if bytes.HasSuffix(buf.Bytes(), []byte(suffix)) && start != 0 {
			buf = bytes.NewBuffer(append(buf.Bytes()[:start]))
			start = 0
		}

	}
}

// Deleted ...
func (b *Note) Deleted() bool {
	return b.Trashed == 1
}

// Marshal ...
func (b *Note) Marshal() ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})

	d, err := yaml.Marshal(b.Frontmatter)
	if err != nil {
		return nil, err
	}

	_, err = buf.WriteString("---\n")
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(d)
	if err != nil {
		return nil, err
	}

	_, err = buf.WriteString("\n---\n")
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(b.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
