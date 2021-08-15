## pongo2-embed-loader

[![Build Status](https://travis-ci.com/nathan-osman/pongo2-embed-loader.svg?branch=main)](https://travis-ci.com/nathan-osman/pongo2-embed-loader)
[![Coverage Status](https://coveralls.io/repos/github/nathan-osman/pongo2-embed-loader/badge.svg?branch=main)](https://coveralls.io/github/nathan-osman/pongo2-embed-loader?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/nathan-osman/pongo2-embed-loader)](https://goreportcard.com/report/github.com/nathan-osman/pongo2-embed-loader)
[![Go Reference](https://pkg.go.dev/badge/github.com/nathan-osman/pongo2-embed-loader.svg)](https://pkg.go.dev/github.com/nathan-osman/pongo2-embed-loader)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

[pongo2](https://github.com/flosch/pongo2) provides Django-like templates for Go applications. However, it does not provide native support for the new [embed directive](https://pkg.go.dev/embed) introduced in Go 1.16. This package aims to fill that gap by providing a `TemplateLoader` for embedded files.

### Usage

The package is extremely easy to use.

Assuming you have a directory named `templates` in the current source directory, you can create a template set with:

```golang
import "github.com/nathan-osman/pongo2-embed-loader"

//go:embed templates/*
var content embed.FS

templateSet := pongo2.NewSet("", &loader.Loader{Content: content})
```

You can then retrieve a template with:

```golang
t, err := templateSet.FromFile("templates/file.tpl")
```
