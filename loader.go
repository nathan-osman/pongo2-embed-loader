package loader

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
)

// Loader implements pongo2's TemplateLoader interface for templates stored
// in a fs.FS (such as an embed.FS).
type Loader struct {
	Content fs.FS
}

// Abs returns the absolute path for the specified template.
func (l *Loader) Abs(base, name string) string {
	return name
}

// Get creates a reader for the specified template.
func (l *Loader) Get(path string) (io.Reader, error) {
	f, err := l.Content.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
