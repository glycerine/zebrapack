package testdata

import "github.com/glycerine/rbtree"

// gotta test the tests too!
//go:generate zebrapack -no-structnames-onwire -unexported

type u struct {
	m map[string]*Tr `zid:"0"`

	// track which job we are up to.
	s string `zid:"1"`
	n int64  `zid:"2"`
}

type Tr struct {
	U  string       `zid:"0"`
	et *rbtree.Tree `msg:"-"`

	Nt int `zid:"1"`

	Snot   map[string]bool `zid:"2"`
	Notyet map[string]bool `zid:"3"`

	Setm []*inn `zid:"4"`
}

type inn struct {
	j map[string]bool `zid:"0"`
	e int64           `zid:"1"`
}
