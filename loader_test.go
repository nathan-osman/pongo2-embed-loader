package loader

import (
	"embed"
	"testing"

	"github.com/flosch/pongo2"
)

var (
	//go:embed test.txt
	content embed.FS

	//go:embed test.txt
	contentStr string
)

func TestLoader(t *testing.T) {
	templateSet := pongo2.NewSet("", &Loader{Content: content})
	tpl, err := templateSet.FromFile("test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	s, err := tpl.Execute(pongo2.Context{})
	if err != nil {
		t.Fatal(err.Error())
	}
	if s != contentStr {
		t.Fatalf("%s != %s", s, contentStr)
	}
}
